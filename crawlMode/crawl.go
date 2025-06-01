package crawlMode

import (
	"bytes"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"xcrawl/fetch"
)

const Reset = "\033[0m"

type LinkInfo struct {
	Path       string
	StatusCode int
	Alive      bool
}

func worker(wg *sync.WaitGroup, parsedURL *url.URL, set mapset.Set[string], l []LinkInfo) {
	defer wg.Done()

	for i := 0; i < len(l); i++ {
		resp, err := http.Get(l[i].Path)
		if err != nil {
			fmt.Printf("Domain is unreachable: %s\n", l[i].Path)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed to read body: %s\n", l[i].Path)
			continue
		}

		doc, err := html.Parse(bytes.NewReader(body))
		if err != nil {
			fmt.Println("HTML parse error:", err)
			continue
		}
		links := ExtractLinks(doc, *parsedURL)
		for _, link := range links {
			if set.Contains(link.Path) {
				continue
			}
			if !set.Contains(link.Path) {
				set.Add(link.Path)
				l = append(l, link)
			}
		}
	}
}

func Run(baseURL string, threads int, depth int) {
	baseURLStatus := fetch.CheckStatuscodeFromURL(baseURL)
	if baseURLStatus != 200 {
		fmt.Printf("url is unreachable %s\n", baseURL)
		os.Exit(1)
	}

	if baseURL[len(baseURL)-1:] != "/" {
		baseURL += "/"
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	set := mapset.NewSet[string]()
	Links := []LinkInfo{
		{
			StatusCode: 200,
			Path:       baseURL,
			Alive:      true,
		},
	}
	var (
		wg sync.WaitGroup
	)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(&wg, parsedURL, set, Links)
	}
	wg.Wait()

	it := set.Iterator()
	for value := range it.C {
		fmt.Printf("Discovered: %s\n", value)
	}
}
