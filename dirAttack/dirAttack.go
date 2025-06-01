package dirAttack

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"xcrawl/fetch"
	"xcrawl/utils"
)

const Reset = "\033[0m"

func worker(jobs <-chan string, results chan<- fetch.Result, wg *sync.WaitGroup, rateLimiter <-chan time.Time) {
	defer wg.Done()
	for url := range jobs {
		<-rateLimiter

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Domain is unreachable %s\n", url)
			continue
		}

		var size int
		if resp.ContentLength == -1 {
			body, _ := io.ReadAll(resp.Body)
			size = len(body)
			_ = resp.Body.Close()
		} else {

			size = int(resp.ContentLength)
			_ = resp.Body.Close()
		}

		res := fetch.Result{
			URL:           url,
			StatusCode:    resp.StatusCode,
			ContentLength: int64(size),
		}
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
	results := make(chan fetch.Result, threads)
	var wg sync.WaitGroup
	rate := time.Second / 5
	rateLimiter := time.NewTicker(rate)

	// var dir string
	if len(baseURL) > 0 && baseURL[len(baseURL)-1] != '/' {
		baseURL = baseURL + "/"
	}

	// Start worker goroutines
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
			res.URL,
			utils.StatusColor(res.StatusCode), // e.g., "\033[32m"
			res.StatusCode,
			Reset, // e.g., "\033[0m"
			res.ContentLength)
	}
}
