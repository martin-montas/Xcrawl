package main

import ( 
	"log"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
)

var links []string

func extractLinks(currDomain string) ([]string, error) {
	// Slice to store extracted href links

	resp, err := http.Get(currDomain)
	if err != nil {
		log.Fatal(err)
		return links, err
	}
	defer resp.Body.Close()

	// TODO:(parse the body here for href hyperlinks of any type of reference to a
	// a page on the same domain and append it to the urls slice)
	// Print the response body

	// Parse the HTML from the response
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Recursive function to traverse the HTML tree
	var extractHrefs func(*html.Node)
	extractHrefs = func(n *html.Node) {

		// If the node is an <a> tag, get the href attribute
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					// process the domain
					ok, fullURL, err := validateDomain(currDomain, attr.Val)
					
					if err == nil {
						log.Fatal(err)
					}

					if ok {
						links = append(links, fullURL)
					}
				}
			}
		}

		// Recursively traverse child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractHrefs(c)
		}
	}

	// Start extracting href attributes from the root of the HTML tree
	extractHrefs(doc)
	return links, nil
}

func validateDomain(baseSite string, foundURI string) (bool, string, error) {
	baseURL, err := url.Parse(baseSite)
	if err != nil {
		return false, "", err
	}

	targetURL, err := url.Parse(foundURI)
	if err != nil {
		return false, "", err
	}

	// Resolve relative URI to absolute
	resolvedURL := baseURL.ResolveReference(targetURL)

	// Compare hostnames
	return resolvedURL.Host == baseURL.Host, resolvedURL.String(), nil
}

