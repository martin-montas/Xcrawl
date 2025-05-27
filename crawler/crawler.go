package crawler

import (
	"sync"
	"xcrawl/fetch"
)

func Run(domain string) {
	var wg sync.WaitGroup
	ch := make(chan fetch.Element)

	rootLink := fetch.Link{
		Path: domain,
	}
	fetch.AppendToLinks(&rootLink)

	wg.Add(1)
	go rootLink.Get(ch, &wg) 

	res := <-ch
	wg.Wait()

	ExtractLinks(*res.Node, *res.Base)

	for _, link := range fetch.GetLinks() {
		wg.Add(1)
		go link.Get(ch, &wg)
		res := <-ch
		wg.Wait()

		ExtractLinks(*res.Node, *res.Base)
	}
	wg.Wait()

}
