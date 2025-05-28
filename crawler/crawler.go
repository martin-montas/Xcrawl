package crawler

import (
	"sync"
	"fmt"
	"time"
	"xcrawl/utils"
	"strings"

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

func Run(domain string, threads int, delay float64) {
	var wg sync.WaitGroup
	ch := make(chan fetch.Element)

	jobs 	:= make(chan string, threads)
	results := make(chan fetch.Result)

	// var dir string
	if !strings.HasSuffix(domain, "/") {
		domain = domain + "/"
	}

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(jobs, results, delay, &wg)
	}

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
	// rootLink := fetch.Link {
	// 	Path: domain,
	// }
	// fetch.AppendToLinks(&rootLink)
	// for _, link := range fetch.GetLinks() {
	// 	wg.Add(1)
	// 	go link.Get(ch, &wg)
	// 	res := <-ch
	// 	wg.Wait()
	// 	ExtractLinks(*res.Node, *res.Base)
	// }
	// wg.Wait()
}
