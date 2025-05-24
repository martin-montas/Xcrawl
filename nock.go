package main

import (
	"fmt"

	"nock/parser"
	"nock/scheduler"

)

func run(domain string, thread int) {
	// utils.PrintInfo("will be using %s threads ", strconv.Itoa(thread))
	fmt.Println(thread)

	// crawls the initial site
	parser.Crawl(domain)

	for _, l := range scheduler.Links {
   	parser.Crawl(l.Path)
	}
}

