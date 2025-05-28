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

var Links []Link

type Link struct {
	Alive      bool
	StatusCode int
	Path       string
	ID         int
}

type Element struct {
	Node *html.Node
	Base *url.URL
	ResponseLength int64  // ResponseLength
}

type  Result struct {
	URL 			string
	Alive      		bool
	StatusCode 		int
	ContentLength 	int64 // ContentLength
}

func (l *Link) Get()  Element {
		resp := fetchAndHandle(l.Path)
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading body from %s:\n", l.Path)
			os.Exit(1)
		}
		n, err := html.Parse(strings.NewReader(string(b)))
		if err != nil {
			log.Printf("Error parsing HTML from %s:\n", l.Path)
			os.Exit(1)
		}
		base, err := url.Parse(l.Path)
		if err != nil {
			log.Printf("Error parsing base URL %s:\n", l.Path)
			os.Exit(1)
		}
		size := resp.ContentLength
		if size == int64(-1) {
			size = int64(3487)
		}

		return  Element {
			Node: n, 
			Base: base, 
			ResponseLength: size,
		}

}

func fetchAndHandle(url string)  *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Domain is unreachable %s\n", url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	return resp
}

func GetStatuscodeFromURL(u string)  Result {
	resp := fetchAndHandle(u) 
	size := resp.ContentLength
	if size == -1 {

		size = 3487
	}
	return Result{
		URL: 			u,
		Alive:      	true,
		StatusCode: 	resp.StatusCode,
		ContentLength: 	size,
	}
}
