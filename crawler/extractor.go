package crawler

import (
	"golang.org/x/net/html"
	"net/url"
	"xcrawl/fetch"
)

func ExtractLinks(doc html.Node, baseUrl url.URL) {
	var (
		tags = []string{
			"a",
			"link",
			"base",
			"area",
		}
	)
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

func processLinks(n html.Node, baseUrl url.URL) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			url, err := url.Parse(attr.Val)
			if err != nil {
				continue
			}
			resolved := baseUrl.ResolveReference(url)
			statusCode := fetch.CheckStatuscodeFromURL(resolved.String())

			l := fetch.Link{
				StatusCode: statusCode,
				Path:       resolved.String(),
			}
			l.DisplayInfo()
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			}
		}
	}
}
