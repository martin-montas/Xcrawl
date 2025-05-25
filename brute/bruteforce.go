package brute

import (
	"bufio"
	"strings"
	"sync"
	"fmt"
	"os"
	"crawlx/request"
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
	ch2 := make(chan int)
	if strings.HasSuffix(domain, "/") {
		domain = domain
	} else {
		domain = domain + "/"
	}
	fmt.Printf("test %s\n", domain)
	for scanner.Scan() {
		wg.Add(1)
		go request.GetLinkStatus(string(domain + scanner.Text()), ch2, &wg)
		res := <-ch2
		wg.Wait()

		if res != 200 {
			continue
		}
		fmt.Printf("\033[32m[%d]\033[0m: %s \n", res, string(domain+scanner.Text()))
		continue
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	wg.Wait()
}
