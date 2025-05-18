package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	//"github.com/go-playground/locales/root"
	"golang.org/x/net/html"
)

var links 			[]string
var elements 	= 	[]string {"a", "base", "area", "link"}
var Nodes  [] html.Node


func sendHTTPRequestAndParse(currDomain string) *html.Node {
	// Slice to store extracted href links
	resp, err := http.Get(currDomain)
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"),
		"[\033[32mINFO\033[0m]", currDomain)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	// TODO:(parse the body here for href hyperlinks of any type of reference to a
	// a page on the same domain and append it to the urls slice)
	// Print the response body

	// Parse the HTML from the response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	// fmt.Println(bodyString)
	doc, err := html.Parse(strings.NewReader(bodyString))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func run(rootDomain string) {
	doc := sendHTTPRequestAndParse(rootDomain)
	extractLinks(*doc)
}

func extractLinks(doc html.Node)   {
    if doc.Type == html.ElementNode && doc.Data == "a" {
        for _, attr := range doc.Attr {
            if attr.Key == "href" {
                fmt.Println("Link:", attr.Val)
            }
        }
        if doc.FirstChild != nil && doc.FirstChild.Type == html.TextNode {
            fmt.Println("Text:", doc.FirstChild.Data)
        }
    }
    for c := doc.FirstChild; c != nil; c = c.NextSibling {
        extractLinks(*c)
    }
}
	// }
// }
func validateDomain(baseSite string, foundURI string) (bool, string, error) {
	baseURL, err := url.Parse(baseSite)
	if err != nil {
		fmt.Println("fuck!!")
		return false, "", err
	}
	targetURL, err := url.Parse(foundURI)
	if err != nil {
		return false, "", err
	}
	// Resolve relative URI to absolute
	resolvedURL := baseURL.ResolveReference(targetURL)

	// Compare hostnames
	fmt.Println(resolvedURL.String())
	return resolvedURL.Host == baseURL.Host, resolvedURL.String(), nil
}

func findDupplicateLinksArray(links []string) []string {
	for ind, link := range links {
		for ind_2, link_2 := range links {
			if ind == ind_2 {
				continue
			} else if link == link_2 {
				removeAtIndex(links, ind_2)
			}
		}
	}
	return links
}

func removeAtIndex(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func catToFile() { }
func pingDomain() { }
