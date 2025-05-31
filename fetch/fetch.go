package fetch

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Path       string
	StatusCode int
	Alive      bool
}

type Element struct {
	URL            string
	Node           *html.Node
	Base           *url.URL
	ResponseLength int64
}

type Result struct {
	URL           string
	StatusCode    int
	ContentLength int64
}

func GetElementFromURL(u string) (Element, error) {
	resp, err := FetchResponse(u)
	if err != nil {
		return Element{}, fmt.Errorf("fetch error for %s: %w", u, err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Element{}, fmt.Errorf("error reading body from %s: %w", u, err)
	}

	n, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		return Element{}, fmt.Errorf("error parsing HTML from %s: %w", u, err)
	}

	base, err := url.Parse(u)
	if err != nil {
		return Element{}, fmt.Errorf("error parsing base URL %s: %w", u, err)
	}

	size := resp.ContentLength
	if size == -1 {
		size = int64(len(b))
	}

	return Element{
		URL:            u,
		Node:           n,
		Base:           base,
		ResponseLength: size,
	}, nil
}

func FetchResponse(url string) (*http.Response, error) {
	if len(url) > 0 && url[len(url)-1] != '/' {
		url += "/"
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("domain is unreachable: %s", url)
	}
	return resp, nil
}

func GetStatuscodeFromURL(u string) Result {
	resp, err := FetchResponse(u)
	if err != nil {
		fmt.Printf("Domain is unreachable %s\n", u)
		os.Exit(1)
	}
	size := resp.ContentLength
	if size == -1 {
		size = 3487
	}
	return Result{
		URL:           u,
		StatusCode:    resp.StatusCode,
		ContentLength: size,
	}
}
