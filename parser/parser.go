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

func Extract(n html.Node, d string) {
	var tags = [4]string{
		"a",
		"link",
		"base",
		"area",
	}
	base, err := url.Parse(d)

	if err != nil {
		fmt.Println("Invalid base URL")
		return
	}
	for _, t := range tags {
		if n.Type == html.ElementNode && n.Data == t {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					u, err := url.Parse(attr.Val)

					if err != nil {
						continue
					}
					Resolved := base.ResolveReference(u)
					Alive, StatusCode := scheduler.IsLinkAlive(Resolved.String())

					link := &worker.Link{
						Alive:      Alive,
						Visited:    true,
						StatusCode: StatusCode,
						Path:       Resolved.String(),
					}
					scheduler.AppendToLink(link)
				}
			}
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				fmt.Println("Text:", n.FirstChild.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Extract(*c, d)
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
