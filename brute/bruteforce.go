package brute

import (
	"bufio"
	"sync"
	"fmt"
	"os"
	"nock/request"
)

func Run(w string, d string) {
	fmt.Println("Starting dictionary attack")
	f, err := os.Open(w)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	wg := sync.WaitGroup{}
	ch2 := make(chan int)
	for scanner.Scan() {
		wg.Add(1)
		go request.GetLinkStatus(string(d + scanner.Text()), ch2, &wg)
		res := <-ch2
		wg.Wait()
		if res != 200 {
			continue
		}
		fmt.Printf("\033[32m[%d]\033[0m: %s \n", res, string(d+scanner.Text()))
		continue
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	wg.Wait()
}
