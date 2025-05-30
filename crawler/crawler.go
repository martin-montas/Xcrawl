package crawler

import (
	"fmt"
	"os"
	"sync"
	"time"

	"xcrawl/fetch"
)

func worker(delay float64, links []fetch.Link, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(links); i++ {
		link := links[i]
		res, err := fetch.GetElementFromURL(link.Path)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		ExtractLinks(*res.Node, *res.Base)
		time.Sleep(time.Duration(delay * float64(time.Second)))
	}
}

func Run(domain string, threads int, delay float64) {
	domainStatus := fetch.CheckStatuscodeFromURL(domain)

	if domainStatus != 200 {
		fmt.Printf("Domain is unreachable %s\n", domain)
		os.Exit(1)
	}

	// var dir string
	if domain[len(domain)-1:] != "/" {
		domain = domain + "/"
	}
	var Links []fetch.Link
	l := fetch.Link{
		Path:       domain,
		StatusCode: domainStatus,
	}
	Links = append(Links, l)

	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(delay, Links, &wg)
	}
	wg.Wait()
}
