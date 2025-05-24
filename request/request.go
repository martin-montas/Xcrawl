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

func processLinks(n html.Node, baseUrl url.URL) {
	for i, attr := range n.Attr {
		if attr.Key == "href" {
			url, err := url.Parse(attr.Val)
			if err != nil {
				continue
			}
			resolved := baseUrl.ResolveReference(url)

			if resolved.Host == baseUrl.Host || resolved.Host == "" {
				alive, statusCode := IsPathAlive(resolved.String())
				l := Link{
					Alive:      alive,
					StatusCode: statusCode,
					Path:       resolved.String(),
					ID:         i,
				}
				l.DisplayInfo()
				AppendToLinks(&l)

			}
		}
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		}
	}
}

func whichSection(n *html.Node) string {
	for p := n.Parent; p != nil; p = p.Parent {
		if p.Type == html.ElementNode {
			if p.Data == "head" {
				return "head"
			}
			if p.Data == "body" {
				return "body"
			}
		}
	}
	return "unknown"
}
