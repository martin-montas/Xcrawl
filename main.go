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
		delay := dirCmd.Float64("d", 0, "Number of threads")

		dirCmd.Parse(os.Args[2:])
		if *url == "" || *wordlist == "" {
			fmt.Println("Usage: dir -u <url> -w <wordlist> -t <threads> -d <delay ms>")
			os.Exit(1)
		}
		utils.InitialInfo(*url, *wordlist, *threads, Version, *delay)
		brute.Run(*wordlist, *url, *threads, *delay)
		os.Exit(0)

	case "crawl":
		crawlCmd := flag.NewFlagSet("crawl", flag.ExitOnError)
		url := crawlCmd.String("u", "", "Target URL")
		threads := crawlCmd.Int("t", 10, "Number of threads")
		delay := crawlCmd.Float64("d", 0.1, "Number of threads")

		crawlCmd.Parse(os.Args[2:])

		if *url == "" {
			fmt.Println("Usage: crawl -u <url> -t <threads> -d <delay ms>")
			os.Exit(1)
		}
		utils.InitialInfo(*url, "", *threads, Version, *delay)
		crawler.Run(*url, *threads, *delay)
		os.Exit(0)
	default:
		fmt.Println("Unknown subcommand:", os.Args[1])
		fmt.Println("Expected 'dir' or 'crawl'")
		os.Exit(1)
	}
}
