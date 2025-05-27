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
	// url 	:= flag.String("u", "", "Initial site")
	// w 		:= flag.String("w", "", "Wordlist")
	// version := flag.Bool("v", false, "Version")
	// t 		:= flag.Int("t", 10, "Threads")
	// flag.Parse()

	// args 	:= flag.Args()

	// if len(args) == 0 {
	// 	fmt.Println("Please specify a command: crawl or dir")
	// 	os.Exit(1)
	// }
	// command := args[0]
	if len(os.Args) < 2 {
		fmt.Println("Expected 'dir' or 'crawl' subcommand")
		os.Exit(1)
	}
	switch os.Args[1] {
		case "dir":
			dirCmd := flag.NewFlagSet("dir", flag.ExitOnError)
			url := dirCmd.String("u", "", "Target URL")
			wordlist := dirCmd.String("w", "", "Wordlist path")
			threads := dirCmd.Int("t", 10, "Number of threads")

			dirCmd.Parse(os.Args[2:])

			if *url == "" || *wordlist == "" {
				fmt.Println("Usage: dir -u <url> -w <wordlist>")
				os.Exit(1)
			}
			utils.InitialInfo(*url, *wordlist, *threads)
			brute.Run(*wordlist, *url)
			os.Exit(0)

		case "crawl":
			crawlCmd := flag.NewFlagSet("crawl", flag.ExitOnError)
			url := crawlCmd.String("u", "", "Target URL")
			threads := crawlCmd.Int("t", 10, "Number of threads")

			crawlCmd.Parse(os.Args[2:])

			if *url == "" {
				fmt.Println("Usage: crawl -u <url>")
				os.Exit(1)
			}

		    utils.InitialInfo(*url, "", *threads)
			crawler.Run(*url, *threads)
			os.Exit(0)

		default:
			fmt.Println("Unknown subcommand:", os.Args[1])
			fmt.Println("Expected 'dir' or 'crawl'")
			os.Exit(1)
		}
}
