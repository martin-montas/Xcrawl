package brute

import (
	"bufio"
	"strings"
	"sync"
	"fmt"
	"os"
	"xcrawl/fetch"
	"xcrawl/utils"
)


const 	Red   = "\033[31m"
const 	Green = "\033[32m"
const 	Blue  = "\033[34m"
const 	Reset = "\033[0m"

func Run(wordlist string, domain string) {
	f, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	wg := sync.WaitGroup{}
	ch2 := make(chan fetch.Status)

	var dir string
	if !strings.HasSuffix(domain, "/") {
		domain = domain + "/"
	}
	for scanner.Scan() {
		wg.Add(1)
		go fetch.GetStatuscodeFromURL(string(domain + scanner.Text()), ch2, &wg)
		res := <-ch2
		wg.Wait()
		if res.StatusCode != 200 {
			continue
		}
		if !strings.HasSuffix(string(scanner.Text()), "/") {
			dir = string(scanner.Text()) + "/"
		}

		color := utils.StatusColor(res.StatusCode)

		fmt.Printf("%-21s %s(Status: %3d)\033[0m\n", dir, color, res.StatusCode)
		continue
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	wg.Wait()
}
