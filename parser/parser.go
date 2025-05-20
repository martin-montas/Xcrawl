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
	"nock/utils"
	"nock/worker"
)

func Extract(node html.Node, domain string) {
	var tags = [4]string{
		"a",
		"link",
		"base",
		"area",
	}
	base, err := url.Parse(domain)

	if err != nil {
		fmt.Println("Invalid base URL")
		return
	}
	id := 0
	for _, tag := range tags {
		if node.Type == html.ElementNode && node.Data == tag {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					id += 1
					url, err := url.Parse(attr.Val)
					if err != nil {
						continue
					}
					resolved := base.ResolveReference(url)
					alive, statusCode := scheduler.IsPathAlive(resolved.String())
					link := &worker.Link{
						Alive:      alive,
						StatusCode: statusCode,
						Path:       resolved.String(),
						ID:         id,
					}
					scheduler.AppendToLink(link)
				}
			}
			if node.FirstChild != nil && node.FirstChild.Type == html.TextNode {
				// fmt.Println("Text:", n.FirstChild.Data)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		Extract(*c, domain)
	}
}

func Parse(n string, t int) *html.Node {
	r, err := http.Get(n)
	utils.PrintInfo(n)

	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	body := string(b)
	d, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	return d
}
