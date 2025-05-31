package crawler

import (
	"fmt"
	"io"
	"bytes"
	// "time"
	"os"
	"sync"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"xcrawl/fetch"
	mapset "github.com/deckarep/golang-set/v2"
)


const Reset = "\033[0m"
func worker(wg *sync.WaitGroup, urls <-chan string, responses chan<- []byte) {
	defer wg.Done()

	for u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			fmt.Printf("4 Domain is unreachable: %s\n", u)
			wg.Done()
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed to read body: %s\n", u)
			wg.Done()
			continue
		}
		responses <- body
		wg.Done()
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

	urls 		:= make(chan string, threads)
	responses 	:= make(chan []byte, threads)
	set 		:= mapset.NewSet[string]()
	var ( 
		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(1)
	urls <- baseURL

	// Start worker goroutines
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(&wg, urls, responses)
	}

	// Start response handler
	go func() {
		for body := range responses {
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
					mu.Lock()
					fmt.Printf("Discovered: %s\n", link.Path)
					mu.Unlock()
					set.Add(link.Path)
					wg.Add(1)
					urls <- link.Path
				}
			}
			wg.Done()
		}
	}()
	wg.Wait()
	close(urls)
	close(responses)
	wg.Done()
}
