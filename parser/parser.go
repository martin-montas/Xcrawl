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

var t int

func Crawl(domain string)  {
	response, err := http.Get(domain)
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

	// return *n, *base
	ExtractLinks(*n, *base)
}

func ExtractLinks(doc html.Node, baseUrl url.URL) {
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
		ExtractLinks(*c, baseUrl)
	}
}

func processLinks(n html.Node, baseUrl url.URL)  {
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
				Node:				n ,
			}
			scheduler.AppendToLink(&l)
	}
	if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		fmt.Println("Text:", n.FirstChild.Data)
		}
	}
}

func GetLinks(s string, thread int) {
}
