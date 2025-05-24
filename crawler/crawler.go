package crawler

import (
	// "fmt"
	"sync"

	"nock/request"
)


// var wg sync.WaitGroup
// jobs := make(chan string, 100) // URLs
// results := make(chan request.Result, 100)
// 
// // Start N workers
// for i := 0; i < 10; i++ {
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for url := range jobs {
// 			result := request.SendBlocking(url) // Blocking version, no channel
// 			results <- result
// 		}
// 	}()
// }
// 
// // Feed jobs
// go func() {
// 	for _, url := range urls {
// 		jobs <- url
// 	}
// 	close(jobs)
// }()
// 
// go func() {
// 	wg.Wait()
// 	close(results)
// }()
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

