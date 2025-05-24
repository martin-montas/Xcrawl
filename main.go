package main

import (
	"flag"
	"fmt"
	"os"

	"nock/brute"
	"nock/crawler"
	"nock/utils"
)

func main() {
	url := flag.String("u", "", "Initial site")
	mode := flag.String("mode", "crawl", "Mode: dir | crawl")
	w := flag.String("w", "", "Wordlist")

	flag.Parse()
	utils.Banner()

	switch *mode {
	case "dir":
		if *url == "" {
			fmt.Println("Please provide a url")
			os.Exit(1)
		}
		if *w == "" {
			fmt.Println("Please provide a wordlist")
			os.Exit(1)
		}
		brute.Run(*w, *url)
		os.Exit(0)

	case "crawl":
		if *url == "" {
			fmt.Println("Please provide a url")
			os.Exit(1)
		}
		fmt.Println("Starting crawler")
		crawler.Run(*url)
		os.Exit(0)

	default:
		fmt.Printf("Unknown mode: %s\n", *mode)
		flag.Usage()
		os.Exit(1)
	}
}
