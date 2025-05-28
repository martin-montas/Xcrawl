package brute


import (
	"sync"
	"time"

	"xcrawl/fetch"
)

func worker(jobs <-chan string, results chan<- fetch.Result, delay float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		res := fetch.GetStatuscodeFromURL(url)
		results <- res
		time.Sleep(time.Duration(delay * float64(time.Second)))
	}
}
