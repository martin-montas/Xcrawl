package main

import (
	"flag"
	"nock/Worker"
	"nock/utils"
)

func main() {
	url := flag.String("u", "", "Root domain/IP")
	version := flag.Bool("v", false, "Enable verbose output.")
	t := flag.Int("t", 3, "The amount of threads.")
	files := flag.String("o", "", "Output file.")
	flag.Parse()
	utils.Banner()
	if *version {
		utils.PrintInfo("Verbose mode is on")
	}
	if !*version {
		utils.PrintInfo("Verbose mode is off")
	}
	utils.PrintInfo("Using threads given")
	if *files != "" {
		utils.PrintInfo("Will be used for saving")
	}
	if *url != "" {
		worker.Run(*url, *t, *version)
	}
	if *url == "" {
		utils.PrintErr("An url should be given")
	}
}
