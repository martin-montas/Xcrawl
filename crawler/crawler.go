package crawler

import (
	"sync"
	"nock/request"
)

func Run(domain string) {
	var wg sync.WaitGroup

	ch := make(chan request.Result)

	wg.Add(1)
	go request.Send(domain, ch, &wg)

	res := <-ch
	wg.Wait()

	ExtractLinks(*res.Node, *res.Base)

	for _, link := range request.GetLinks() {
		wg.Add(1)
		go request.Send(link.Path, ch, &wg)
		res := <-ch
		wg.Wait()

		ExtractLinks(*res.Node, *res.Base)
	}
	wg.Wait()

}
