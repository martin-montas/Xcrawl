package main

import (
	"nock/parser"
	"nock/utils"
)

func run(domain string, thread int, verbose bool) {
	utils.PrintInfo("using the threads given")
	if verbose {
		utils.PrintInfo("verbose set to true")

		for {
			parser.Crawl(domain)
		}

	} else {
		utils.PrintInfo("verbose set to false")

	}

}
