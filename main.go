package main

import (
	"flag"
	"fmt"
	"os"

	"xcrawl/brute"
	"xcrawl/crawler"
	"xcrawl/utils"
)

const (
	Version = "1.5.0"
)

func main() {
	url 	:= flag.String("u", "", "Initial site")
	dir 	:= flag.Bool("dir", false, "dir finding attack")
	crawl 	:= flag.Bool("crawl", false, "crawler")
	w 		:= flag.String("w", "", "Wordlist")
	version := flag.Bool("v", false, "Version")
	t 		:= flag.Int("t", 10, "Threads")

	flag.Parse()
	utils.Banner(Version)

	if *version {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	if *dir {
		if *url == "" {
			fmt.Println("Please provide a url")
			os.Exit(1)
		}
		if *w == "" {
			fmt.Println("Please provide a wordlist")
			os.Exit(1)
		}
		utils.InitialInfo(*url, *w,*t)
		brute.Run(*w, *url)
		os.Exit(0)
	}
    if *crawl {
		if *url == "" {
			fmt.Println("Please provide a url")
			os.Exit(1)
		}
		utils.InitialInfo(*url, *w,*t)
		crawler.Run(*url)
		os.Exit(0)

	}
}
