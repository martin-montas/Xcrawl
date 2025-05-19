package crawl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

var t = []string{"a", "base", "area", "link"}

func Run(d string) {
	n := parse(d)
	extract(*n, d)
}

func parse(n string) *html.Node {
	r, err := http.Get(n)
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"),
		"[\033[32mINFO\033[0m]", n)

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

func extract(n html.Node, d string) {
	var links [] Link
	for _, t := range t {
		if n.Type == html.ElementNode && n.Data == t {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println("Link:", attr.Val)
					links = append(links, Link{path: attr.Val, text: attr.Val})
				}
			}
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				fmt.Println("Text:", n.FirstChild.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extract(*c, d)
	}
}
