package cmd

import (
	"fmt"
	"os"

	"nock/nockdir"
)

var (
	Version     = "1.0.0"
	defaultList = "/usr/share/wordlists/dirb/small.txt"
)

type Parser interface {
	Parse(args []string, version string)
}

// modules should be call here
var registry = map[string]Parser{
	"dir": &nockdir.NockDir{},
}

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

func Execute() {
	command := os.Args[1]
	if command == "version" || command == "-v" || command == "--version" {
		fmt.Println("Nock", Version)
		return
	}
	if command == "" {
		fmt.Println("Expected 'dir' , 'crawl' or 'version' subcommand")
		return
	}
	d, ok := registry[command]
	if !ok {
		fmt.Println("Unknown command:", command)
		return
	}
	Parse(d)
}

func Parse(c Parser) {
	c.Parse(os.Args[2:], Version)
}
