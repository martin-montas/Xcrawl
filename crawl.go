package main

import ( 
	"log"
	"golang.org/x/net/html"
	"net/http"
)

func extractLinks(urls []string) ([]string, error) {
	// Slice to store extracted href links
	var links []string

	resp, err := http.Get(urls[0])
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
					links = append(links, attr.Val)
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
