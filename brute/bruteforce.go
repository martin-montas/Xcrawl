package brute

import (
	"bufio"
	"strings"
	"sync"
	"fmt"
	"os"
	"time"

	"xcrawl/fetch"
	"xcrawl/utils"
)

const 	Red   = "\033[31m"
const 	Green = "\033[32m"
const 	Blue  = "\033[34m"
const 	Reset = "\033[0m"


func worker(jobs <-chan string, results chan<- fetch.Result, delay float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		res := fetch.GetStatuscodeFromURL(url)
		results <- res
		time.Sleep(time.Duration(delay * float64(time.Second)))
	}
}
func Run(wordlist string, domain string, threads int, delay float64) {
	f, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	jobs 	:= make(chan string, threads)
	results := make(chan fetch.Result)
	var 	wg sync.WaitGroup

	// var dir string
	if !strings.HasSuffix(domain, "/") {
		domain = domain + "/"
	}
	
	// Start worker goroutines
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(jobs, results, delay, &wg)
	}

	// Read the wordlist and enqueue jobs
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			path := scanner.Text()
			url := domain + path
			jobs <- url
		}
		close(jobs)
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for res := range results {
		if res.StatusCode != 200  {
			continue
		}
		dir := res.URL

		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
		parts := strings.SplitN(dir, "/", 3)

		if len(parts) < 3 {
			fmt.Println("The path does not contain two slashes.")
			return
		}

		// Reconstruct the string starting from the second slash
		result 		:= parts[1] + "/" + parts[2]
		statusColor := utils.StatusColor(res.StatusCode)
		resetColor 	:= "\033[0m"

		fmt.Printf("%-30s  %-7s(Status: %-3d)%-7s  [Size: %-4d]\n",
			result,
			statusColor,
			res.StatusCode,
			resetColor,
			res.ContentLength)
	}
}
