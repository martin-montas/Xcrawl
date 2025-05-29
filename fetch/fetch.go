package fetch

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	StatusCode 		int
	Path       		string
}

type Element struct {
	URL 			string
	Node 			*html.Node
	Base 			*url.URL
	ResponseLength int64
}

type  Result struct {
	URL 			string
	StatusCode 		int
	ContentLength 	int64 
}

func  GetElementfromURL(u string)  Element {
		resp   := FetchResponse(u)
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading body from %s:\n", u)
			os.Exit(1)
		}
		n, err := html.Parse(strings.NewReader(string(b)))
		if err != nil {
			log.Printf("Error parsing HTML from %s:\n", u)
			os.Exit(1)
		}
		base, err := url.Parse(u)
		if err != nil {
			log.Printf("Error parsing base URL %s:\n", u)
			os.Exit(1)
		}
		size := resp.ContentLength
		if size == int64(-1) {
			size = int64(3487)
		}

		return Element {
			URL:			u,
			Node: 			n, 
			Base: 			base, 
			ResponseLength: size,
		}
}

func FetchResponse(url string)  *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Domain is unreachable %s\n", url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	return resp
}

func GetStatuscodeFromURL(u string)  Result {
	resp := FetchResponse(u) 
	size := resp.ContentLength
	if size == -1 {
		size = 3487
	}
	return Result{
		URL: 			u,
		StatusCode: 	resp.StatusCode,
		ContentLength: 	size,
	}
}
