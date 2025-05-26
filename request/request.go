package request

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"os"

	"golang.org/x/net/html"
)

var Links []Link

type Link struct {
	Alive      bool
	StatusCode int
	Path       string
	ID         int
}

type Tag struct {
	Node *html.Node
	Base *url.URL
}

type Status struct {
	Alive      bool
	StatusCode int
}

func Send(domain string, ch chan Tag, wg *sync.WaitGroup) {
	fetchAndHandle(domain, wg, func(resp *http.Response) {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading body from %s:\n", domain)
			os.Exit(1)
		}
		n, err := html.Parse(strings.NewReader(string(b)))
		if err != nil {
			log.Printf("Error parsing HTML from %s:\n", domain)
			os.Exit(1)
		}
		base, err := url.Parse(domain)
		if err != nil {
			log.Printf("Error parsing base URL %s:\n", domain)
			os.Exit(1)
		}
		ch <- Tag{Node: n, Base: base}
	})
}

func fetchAndHandle(url string, wg *sync.WaitGroup, handler func(*http.Response)) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Domain is unreachable %s\n", url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	handler(resp)
}

func GetStatuscodeFromURL(u string, ch chan Status, wg *sync.WaitGroup) {
	fetchAndHandle(u, wg, func(resp *http.Response) {
		ch <- Status{Alive: true, StatusCode: resp.StatusCode}
	})
}
