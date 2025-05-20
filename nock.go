package main

import (
	"nock/parser"
)

func run(d string, t int, v bool) {
	n := parser.Parse(d)
	if v {
		parser.Extract(*n, d)
	}
}
