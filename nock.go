package main
import (

	"nock/parser"
	"nock/scheduler"
	"sync"

)

func run(domain string) {
	// crawls the initial site
	var wg sync.WaitGroup
	wg.Add(1)
	go parser.Crawl(domain, &wg)
	wg.Wait()

	for _, l := range scheduler.Links {
		wg.Add(1)
   	go parser.Crawl(l.Path, &wg)
	}
	wg.Wait()
}

