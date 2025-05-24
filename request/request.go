package request

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Link struct {
	Alive      bool
	StatusCode int
	Path       string
	ID         int
}

type Tag struct {
	Node *html.Node
	Base *url.URL
}

type Status struct {
	Alive      bool
	StatusCode int
}

var Links []Link

func  Send(domain string, ch chan Tag, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(domain)

	if err != nil {
		log.Printf("Error fetching %s: %v\n", domain, err)
		return
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading body from %s: %v\n", domain, err)
		return
	}
	body := string(b)
	n, err := html.Parse(strings.NewReader(body))
	if err != nil {
		log.Printf("Error parsing HTML from %s: %v\n", domain, err)
		return
	}

	base, err := url.Parse(domain)
	if err != nil {
		log.Printf("Error parsing base URL %s: %v\n", domain, err)
		return
	}
	ch <- Tag{Node: n, Base: base}
}

func GetStatuscodeFromURL(u string, ch chan Status, wg *sync.WaitGroup)  {
	defer wg.Done()
	response, err := http.Get(u)
	if err != nil {
		fmt.Printf("Domain is unreachable %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		ch <- Status{Alive: true, StatusCode: response.StatusCode}
	}

	ch <- Status{Alive: true, StatusCode: response.StatusCode}
}	

func GetLinkStatus(u string, ch chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	response, err := http.Get(u)
	if err != nil {
		fmt.Printf("Domain is unreachable %s", err)
	}
	defer response.Body.Close()

	 ch <-response.StatusCode
}
