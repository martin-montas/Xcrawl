package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"nock/scheduler"
	"nock/worker"
	"nock/utils"
)

func parse(domain string, thread int) (html.Node , url.URL) {
	response, err := http.Get(domain)

	fmt.Printf("test: %d", thread)
	utils.PrintInfo("querying the domain")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(b)
	n, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	base, err := url.Parse(domain)
	if err != nil {
		log.Fatal(err)
	}

	return *n, *base
}

func extractLinks(doc html.Node, baseUrl url.URL) {
	var tags = []string{
		"a",
		"link",
		"base",
		"area",
	}
	for _, tag := range tags {
		if doc.Type != html.ElementNode || doc.Data != tag {
			continue
		}
		processLinks(doc, baseUrl)

	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(*c, baseUrl)
	}
}

func processLinks(n html.Node, baseUrl url.URL) {
	for i, attr := range n.Attr {
		if attr.Key == "href" {
			// fmt.Println("Link:", attr.Val)
			url, err := url.Parse(attr.Val)
			if err != nil {
				continue
			}
			resolved := baseUrl.ResolveReference(url)
			alive, statusCode := scheduler.IsPathAlive(resolved.String())

			l := worker.Link {
				Alive:      	alive,
				StatusCode:  	statusCode,
				Path:       	resolved.String(),
				ID: 				i,
			}
			worker.Links = append(worker.Links, l)
	}
	if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		fmt.Println("Text:", n.FirstChild.Data)
		}
	}
}

func GetLinks(s string, thread int) {
	n, baseUrl := parse(s, thread)
	extractLinks(n, baseUrl)
}
