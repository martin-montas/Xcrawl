package parser

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"golang.org/x/net/html"

	"nock/utils"
	"nock/worker"
	"nock/scheduler"
)

var tags = [4]string {
	"a",
	"link",
	"base",
	"area",
}

func Extract(n html.Node, d string) {
	for _, t := range tags {
		if n.Type == html.ElementNode && n.Data == t {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					utils.FullURL()


					// scheduler.IsLinkAlive()


					link := worker.Link {
						text: 	attr.Val,
						visited: true,
						parent: 	*n.Parent,
					}
				}
			}
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				// fmt.Println("Text:", n.FirstChild.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Extract(*c, d)
	}
}

func Parse(n string) *html.Node {
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

