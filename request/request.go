package request

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Result struct {
	Node *html.Node
	Base *url.URL
}

func Send(domain string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(domain)

	if err != nil {
		log.Printf("Error fetching %s: %v\n", domain, err)
		return
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading body from %s: %v\n", domain, err)
		return
	}
	body := string(b)
	n, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Printf("Error parsing HTML from %s: %v\n", domain, err)
		return
	}

	base, err := url.Parse(domain)
	if err != nil {
		log.Printf("Error parsing base URL %s: %v\n", domain, err)
		return
	}

	ch <- Result{Node: n, Base: base}
}


