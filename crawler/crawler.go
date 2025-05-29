package crawler

import (
	"sync"
	"fmt"
	"os"
	"time"

	"xcrawl/utils"
	"xcrawl/fetch"
)

func worker(result <-chan []fetch.Link, delay float64, links []fetch.Link, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < len(links); i++ {
		link 	:= links[i]
		res 	:= fetch.GetElementfromURL(link.Path)
		newLink := ExtractLinks(*res.Node, *res.Base)
		links    = append(links, newLink)
		time.Sleep(time.Duration(delay * float64(time.Second)))
	}
	result <- links
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
	l 	:= fetch.Link {
		Path: domain,
		StatusCode: domainStatus,
	}
	Links 	 = append(Links, l)

	// jobs 	:= make(chan string, threads)
	// results := make(chan []fetch.Link, threads)
	ch := make(chan []fetch.Link)
	

	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(ch, delay, Links, &wg)
		linkSlice := <- ch 
	}

	wg.Wait()
	//	go func() {
	//		for _, l := range Links {
	//			jobs <- l.Path
	//		}
	//		close(jobs)
	//}()
	//
	//	go func() {
	//		wg.Wait()
	//		close(results)
	//	}()
	//
	go func() {
		for l := range results {
			fmt.Printf("%-30s  %-7s(Status: %-3d)%-7s  [Size: %-4d]\n",
				l.Path,
				utils.Green,
				l.StatusCode,
				utils.Reset,
				l.Size)
		}
		close(results)
	}()
}
