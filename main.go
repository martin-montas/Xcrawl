package main

import (
	"flag"

	"nock/utils"
)

func main() {
	url := flag.String("u", "", "Root domain/IP")
	files := flag.String("o", "", "Output file.")
	flag.Parse()
	utils.Banner()
	// utils.PrintInfo("Using threads given")

	if *files != "" {
		utils.PrintInfo("nothing will be used for saving", "")
	}
	if *url != "" {
		run(*url)
	} else {
		utils.PrintErr("An url should be given", "")
	}
}
