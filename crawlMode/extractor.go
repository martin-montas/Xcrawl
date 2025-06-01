package crawlMode

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"xcrawl/fetch"
)

var tags = []string{"a", "link", "base", "area"}

func ExtractLinksFromNode(n *html.Node, baseURL url.URL) []LinkInfo {
	var links []LinkInfo
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			parsed, err := url.Parse(attr.Val)
			if err != nil {
				fmt.Println("Bad URL:", err)
				continue
			}
			resolved := baseURL.ResolveReference(parsed)
			response, err := fetch.FetchResponse(resolved.String())
			if err != nil {
				continue
			}
			if SameDomain(baseURL.String(), resolved.String()) && response.StatusCode == 200 {
				links = append(links, LinkInfo{
					StatusCode: response.StatusCode,
					Path:       resolved.String(),
					Alive:      true,
				})
			} else {
				continue
			}
		}
	}
	return links
}

func extractRecursive(doc *html.Node, baseUrl url.URL, links []LinkInfo) []LinkInfo {
	if doc.Type == html.ElementNode && contains(tags, doc.Data) {
		if whichSection(doc) != "head" {
			newLinks := ExtractLinksFromNode(doc, baseUrl)
			for _, link := range newLinks {
				if link.Alive {
					links = append(links, link)
				}
			}
		}
	}
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

func ExtractLinks(doc *html.Node, baseUrl url.URL) []LinkInfo {
	var links []LinkInfo
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
