package main

import (
	"nock/parser"
	// "nock/utils"
)

func run(domain string, thread int, verbose bool) {
	if verbose {
		parser.GetLinks(domain, thread)
	}
}
