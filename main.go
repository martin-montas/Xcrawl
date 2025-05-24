package main

import (
	"flag"

	"nock/utils"
)

func main() {
	url := flag.String("u", "", "Root domain/IP")
	thread := flag.Int("t", 3, "The amount of threads.")
	files := flag.String("o", "", "Output file.")
	flag.Parse()
	utils.Banner()
	// utils.PrintInfo("Using threads given")

	if *files != "" {
		utils.PrintInfo("nothing will be used for saving", "")
	}
	if *url != "" {
		run(*url, *thread)
	} else {
		utils.PrintErr("An url should be given", "")
	}
}
