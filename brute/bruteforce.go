package brute

import (
	"bufio"
	"strings"
	"sync"
	"fmt"
	"os"
	"time"
	"xcrawl/request"
)

func Run(wordlist string, domain string) {
	f, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	wg := sync.WaitGroup{}
	ch2 := make(chan request.Status)
	if !strings.HasSuffix(domain, "/") {
		domain = domain + "/"
	}
	fmt.Printf("test %s\n", domain)
	for scanner.Scan() {
		wg.Add(1)
		go request.GetStatuscodeFromURL(string(domain + scanner.Text()), ch2, &wg)
		res := <-ch2
		wg.Wait()

		if res.StatusCode != 200 {
			continue
		}
	fmt.Printf("%s \033[32m[%d]\033[0m: %s \n", time.Now().Format("2006-01-02 03:04:05 PM"), res.StatusCode, string(domain+scanner.Text()))
		continue
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	wg.Wait()
}
