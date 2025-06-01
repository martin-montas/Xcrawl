package nockdir

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"nock/httputils"
	utils "nock/ioutils"
)

const Reset = "\033[0m"

type NockDir struct {
	options *OptionsDir
	client  *response.HTTPClient
}

const defaultList = "/usr/share/dirb/wordlists/dirb/common.txt"

func (d *NockDir) Parse(args []string, version string) {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'dir' , 'crawl' or 'version' subcommand")
		return
	}

	dirCmd := flag.NewFlagSet("dir", flag.ExitOnError)
	u := dirCmd.String("u", "", "Target URL")
	w := dirCmd.String("w", defaultList, "Wordlist path")
	t := dirCmd.Int("t", 10, "Number of threads")

	if err := dirCmd.Parse(os.Args[2:]); err != nil {
		log.Fatalf("failed to parse dir command: %v", err)
	}

	if *u == "" || *w == "" {
		fmt.Println("Usage: dir -u <url> -w <wordlist> -t <threads>")
		os.Exit(1)
	}

	// for debugging:
	// fmt.Printf("URL: %s\n", *u)

	opts := &OptionsDir{
		Wordlist: *w,
		BaseURL:  *u,
		Threads:  *t,
	}
	d.options = opts
}

func GetResponseData(resp *http.Response) response.ResponseData {
	return response.ResponseData{
		Name:          resp.Request.URL.String(),
		StatusCode:    resp.StatusCode,
		ContentLength: resp.ContentLength,
	}
}

func (d *NockDir) worker(jobs <-chan string, results chan<- response.ResponseData, wg *sync.WaitGroup, rateLimiter <-chan time.Time) {
	defer wg.Done()
	for path := range jobs {
		<-rateLimiter
		resp, err := d.client.Get(path)
		if err != nil {
			continue
		}
		getResponseData := GetResponseData(resp)

		results <- getResponseData
	}
}

func (d *NockDir) Run() {
	f, err := os.Open(d.options.Wordlist)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	jobs := make(chan string, d.options.Threads)
	results := make(chan response.ResponseData, d.options.Threads)
	var wg sync.WaitGroup
	rate := time.Second / 5

	rateLimiter := time.NewTicker(rate)
	if len(d.options.BaseURL) > 0 && d.options.BaseURL[len(d.options.BaseURL)-1] != '/' {
		d.options.BaseURL += "/"
	}
	for i := 0; i < d.options.Threads; i++ {
		wg.Add(1)
		go d.worker(jobs, results, &wg, rateLimiter.C)
	}
	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			path := scanner.Text()
			url := d.options.BaseURL + path
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
			utils.StatusColor(res.StatusCode),
			res.StatusCode,
			Reset,
			res.ContentLength)
	}
}
