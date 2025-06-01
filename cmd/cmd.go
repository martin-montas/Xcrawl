package cmd

import (
	"fmt"
	"os"

	"nock/nockdir"
)

var Version = "1.0.0"

type Parser interface {
	Parse(version string)
}

// modules should be call here
var registry = map[string]Parser{
	"dir": &nockdir.NockDir{},
}

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
	c.Parse(Version)
}
