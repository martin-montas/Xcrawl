package dirforcer

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"xcrawl/httputils"
	utils "xcrawl/ioutils"
)

const Reset = "\033[0m"

func worker(jobs <-chan string, results chan<- httputils.URL, wg *sync.WaitGroup, rateLimiter <-chan time.Time) {
	defer wg.Done()
	for url := range jobs {
		<-rateLimiter
		res := httputils.NewURL(url)
		results <- res
	}
}

func Run(wordlist string, baseURL string, threads int) {
	f, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	jobs := make(chan string, threads)
	results := make(chan httputils.URL, threads)
	var wg sync.WaitGroup
	rate := time.Second / 5

	rateLimiter := time.NewTicker(rate)
	if len(baseURL) > 0 && baseURL[len(baseURL)-1] != '/' {
		baseURL += "/"
	}
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg, rateLimiter.C)
	}
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			path := scanner.Text()
			url := baseURL + path
			jobs <- url
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
		close(jobs)
	}()

	// Collect results in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		if utils.IsForbidden(res.StatusCode) {
			continue
		}
		fmt.Printf("%-40s %sStatus: %3d%s [Size: %5d]\n",
			res.Name,
			utils.StatusColor(res.StatusCode), // e.g., "\033[32m"
			res.StatusCode,
			Reset, // e.g., "\033[0m"
			res.ContentLength)
	}
}
