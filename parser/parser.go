package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	// "net/url"
	"strings"

	// "nock/scheduler"
	// "nock/worker"
	"nock/utils"
)

func parse(domain string, thread int) html.Node {
	resp, err := http.Get(domain)
	fmt.Printf("%d", thread)
	utils.PrintInfo("querying the domain")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(b)
	n, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	return *n
}

func extractLinks(doc html.Node) {
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
		processLinks(doc)

	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(*c)
	}
}

func processLinks(doc html.Node) {
	for _, attr := range doc.Attr {
		if attr.Key == "href" {
			fmt.Println("Link:", attr.Val)
		}
	}
	if doc.FirstChild != nil && doc.FirstChild.Type == html.TextNode {
		fmt.Println("Text:", doc.FirstChild.Data)
	}
}

func GetLinks(s string, thread int) {
	n := parse(s, thread)
	extractLinks(n)
}
