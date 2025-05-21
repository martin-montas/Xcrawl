package main

import (
	"nock/parser"
	"nock/utils"
)

func run(domain string, thread int, verbose bool) {
	if verbose {
		utils.PrintInfo("verbose set to true")
		parser.GetLinks(domain, thread)
	}
}
