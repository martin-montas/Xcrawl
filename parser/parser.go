package parser

import (
	"io"
	"sync"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"nock/scheduler"
)


func Crawl(domain string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(domain)

	if err != nil {
		log.Printf("Error fetching %s: %v\n", domain, err)
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

	ExtractLinks(*n,*base) 
}


func ExtractLinks(doc html.Node, baseUrl url.URL) {
	var tags = []string {
		"a",
		"link",
		"base",
		"area",
	}
	for _, tag := range tags {
		if doc.Type != html.ElementNode || doc.Data != tag {
			continue
		}
		if whichSection(&doc) == "head" {
			continue

		}
		processLinks(doc, baseUrl)

	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		ExtractLinks(*c, baseUrl)
	}
}

func whichSection(n *html.Node) string {
	// returns if its a head html.Node value
	// or the body 
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

func processLinks(n html.Node, baseUrl url.URL)  {
	for i, attr := range n.Attr {
		if attr.Key == "href" {
			url, err := url.Parse(attr.Val)
			if err != nil {
				continue
			}
			resolved 			:= baseUrl.ResolveReference(url)

			if resolved.Host == baseUrl.Host || resolved.Host == "" {
				alive, statusCode := scheduler.IsPathAlive(resolved.String())
				l := scheduler.Link {
					Alive:      	alive,
					StatusCode:  	statusCode,
					Path:       	resolved.String(),
					ID: 				i,
				}
				l.DisplayInfo()
				scheduler.AppendToLinks(&l)

			} 
		}
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			// fmt.Println("Text:", n.FirstChild.Data)
		}
	}
}

