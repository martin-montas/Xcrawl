package main

import (
	"flag"
	"log"
	"fmt"
	"os"

	"xcrawl/brute"
	"xcrawl/crawler"
	"xcrawl/utils"
)

const (
	Version     = "1.0.0"
	defaultList = "/usr/share/wordlists/dirb/small.txt"
)

func main() {
	if os.Args[1] == "version" || os.Args[1] == "-v" {
		fmt.Println("Version:", Version)
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("Expected 'dir' , 'crawl' or 'version' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dir":
		dirCmd := flag.NewFlagSet("dir", flag.ExitOnError)
		url := dirCmd.String("u", "", "Target URL")
		wordlist := dirCmd.String("w", defaultList, "Wordlist path")
		threads := dirCmd.Int("t", 10, "Number of threads")

		if err := dirCmd.Parse(os.Args[2:]); err != nil {
			log.Fatalf("failed to parse dir command: %v", err)
		}

		if *url == "" || *wordlist == "" {
			fmt.Println("Usage: dir -u <url> -w <wordlist> -t <threads>")
			os.Exit(1)
		}
		utils.InitialInfo(*url, *wordlist, *threads, Version)
		brute.Run(*wordlist, *url, *threads)
		os.Exit(0)

	case "crawl":
		crawlCmd := flag.NewFlagSet("crawl", flag.ExitOnError)
		url := crawlCmd.String("u", "", "Target URL")
		threads := crawlCmd.Int("t", 10, "Number of threads")
		depth := crawlCmd.Int("d", 1, "depth of the crawler")

		if err := crawlCmd.Parse(os.Args[2:]); err != nil {
			log.Fatalf("failed to parse dir command: %v", err)
		}
		if *url == "" {
			fmt.Println("Usage: crawl -u <url> -t <threads> -d <depth>")
			os.Exit(1)
		}
		utils.InitialInfo(*url, "", *threads, Version)
		crawler.Run(*url, *threads, *depth)
		os.Exit(0)
	default:
		fmt.Println("Unknown subcommand:", os.Args[1])
		fmt.Println("Expected 'dir' or 'crawl'")
		os.Exit(1)
	}
}
