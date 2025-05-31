package crawler

import (
	"fmt"
	"io"
	"bytes"
	"time"
	"os"
	"sync"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"xcrawl/fetch"
)

// var m  sync.Mutex
//m.Unlock()
const Reset = "\033[0m"
var client  = &http.Client{ Timeout: 5 * time.Second }

func worker(wg *sync.WaitGroup, urls <-chan string, responses chan<- []byte) {
	defer wg.Done()

	for u := range urls {
		resp, err := client.Get(u)
		if err != nil {
			fmt.Printf("Domain is unreachable: %s\n", u)
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed to read body: %s\n", u)
			continue
		}
		responses <- body
	}
}

func Run(baseURL string, threads int, delay float64) {
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

	urls := make(chan string, threads)
	responses := make(chan []byte, threads)
	var wg sync.WaitGroup

	seen := make(map[string]bool)
	var mu sync.Mutex

	// Start worker goroutines
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(&wg, urls, responses)
	}

	// Seed with baseURL
	go func() {
		urls <- baseURL
	}()

	// Response handler
	go func() {
		for body := range responses {
			doc, err := html.Parse(bytes.NewReader(body))
			if err != nil {
				fmt.Println("HTML parse error:", err)
				continue
			}
			links := ExtractLinks(doc, *parsedURL)

			for _, link := range links {
				mu.Lock()
				fmt.Println("Discovered:", link.Path)
				if !seen[link.Path] {
					seen[link.Path] = true
					urls <- link.Path
				}
				mu.Unlock()
			}
		}
	}()

	// Wait for all workers, then close responses
	go func() {
		wg.Wait()
		close(responses)
	}()

	// Wait for a while to allow crawling (simplified)
	time.Sleep(10 * time.Second)
	close(urls)

	wg.Done()
}
