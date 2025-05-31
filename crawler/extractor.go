package crawler

import (
	"fmt"
	"strings"
	"net/url"

	"golang.org/x/net/html"
	"xcrawl/fetch"
)

var tags = []string{"a", "link", "base", "area"}

func ExtractLinks(doc *html.Node, baseUrl url.URL) []fetch.Link {
	var links []fetch.Link
	return extractRecursive(doc, baseUrl, links)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func extractRecursive(doc *html.Node, baseUrl url.URL, links []fetch.Link) []fetch.Link {
	// If the node matches a tag and isn't in the <head>, attempt to extract link
	if doc.Type == html.ElementNode && contains(tags, doc.Data) {
		if whichSection(doc) != "head" {
			link := ExtractLinksFromNode(doc, baseUrl)
			if link.Alive {
				links = append(links, *link)
			}
		}
	}

	// Recursively search all child nodes
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		links = extractRecursive(c, baseUrl, links)
	}

	return links
}

func whichSection(n *html.Node) string {
	// Determine if node is in <head> or <body>
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

// SameDomain checks if urlA and urlB are in the same domain or subdomain relationship
func SameDomain(urlA, urlB string) bool {
	parsedA, err1 := url.Parse(urlA)
	parsedB, err2 := url.Parse(urlB)
	if err1 != nil || err2 != nil {
		return false
	}
	hostA := strings.ToLower(parsedA.Hostname())
	hostB := strings.ToLower(parsedB.Hostname())

	return hostA == hostB || strings.HasSuffix(hostA, "."+hostB) || strings.HasSuffix(hostB, "."+hostA)
}

func ExtractLinksFromNode(n *html.Node, baseURL url.URL) *fetch.Link {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			parsed, err := url.Parse(attr.Val)
			if err != nil {
				fmt.Println("Bad URL:", err)
				return &fetch.Link{Alive: false}
			}
			resolved := baseURL.ResolveReference(parsed)
			statusCode := fetch.CheckStatuscodeFromURL(resolved.String())
			if SameDomain(baseURL.String(), resolved.String()) && statusCode == 200 {
				return &fetch.Link{
					StatusCode: statusCode,
					Path:       resolved.String(),
					Alive:      true,
				}
			}
		}
	}
	return &fetch.Link{
		StatusCode: 0,
		Path:       "",
		Alive:      false,
	}
}
