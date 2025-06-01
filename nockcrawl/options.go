package nockcrawl

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func crawlerCmd() {
	crawlCmd := flag.NewFlagSet("crawl", flag.ExitOnError)
	url := crawlCmd.String("u", "", "Target URL")
	// threads := crawlCmd.Int("t", 10, "Number of threads")
	if err := crawlCmd.Parse(os.Args[2:]); err != nil {
		log.Fatalf("failed to parse dir command: %v", err)
	}
	if *url == "" {
		fmt.Println("Usage: crawl -u <url> -t <threads>")
		os.Exit(1)
	}
	// utils.InitialInfoCrawler(*url, *threads, Version)
	// crawler.Run(*url, *threads)
	os.Exit(0)
}
