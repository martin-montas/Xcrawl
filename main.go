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
	mode 	:= flag.String("mode", "crawl", "Mode: dir | crawl")
	w 		:= flag.String("w", "", "Wordlist")
	version := flag.Bool("v", false, "Version")


	flag.Parse()
	utils.Banner()

	if *version {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}


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
