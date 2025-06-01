package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	utils "xcrawl/ioutils"
	"xcrawl/nockcrawl"
	"xcrawl/nockdir"
)

const (
	Version     = "1.0.0"
	defaultList = "/usr/share/wordlists/dirb/small.txt"
)

func crawlerCmd() {
	crawlCmd := flag.NewFlagSet("crawl", flag.ExitOnError)
	url := crawlCmd.String("u", "", "Target URL")
	threads := crawlCmd.Int("t", 10, "Number of threads")
	// depth := crawlCmd.Int("d", 1, "depth of the crawler")

	if err := crawlCmd.Parse(os.Args[2:]); err != nil {
		log.Fatalf("failed to parse dir command: %v", err)
	}
	if *url == "" {
		fmt.Println("Usage: crawl -u <url> -t <threads>")
		os.Exit(1)
	}
	utils.InitialInfoCrawler(*url, *threads, Version)
	crawler.Run(*url, *threads)
	os.Exit(0)
}

func dirCmd() {
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
	utils.InitialInfoDirForcer(*url, *wordlist, *threads, Version)
	dirforcer.Run(*wordlist, *url, *threads)
	os.Exit(0)
}

func main() {
	if os.Args[1] == "version" || os.Args[1] == "-v" {
		fmt.Println("Version:", Version)
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		fmt.Println("Expected 'dir' , 'crawl' or 'version' subcommand")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "dir":
		dirCmd()
		os.Exit(0)

	case "crawl":
		crawlerCmd()
		os.Exit(0)
	default:
		fmt.Println("Unknown subcommand:", os.Args[1])
		fmt.Println("Expected 'dir' or 'crawl'")
		os.Exit(1)
	}
}
