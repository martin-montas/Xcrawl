package main

import (
	"flag"
	"fmt"
	"nock/utils"
)

func main() {
	url := flag.String("u", "", "Root domain/IP")
	verbose := flag.Bool("v", true, "Enable verbose output.")
	thread := flag.Int("t", 3, "The amount of threads.")
	files := flag.String("o", "", "Output file.")
	flag.Parse()
	utils.Banner()
	if *verbose {
		utils.PrintInfo("Verbose mode is on")
	}
	if !*verbose {
		utils.PrintInfo("Verbose mode is off")
	}
	utils.PrintInfo("Using threads given")
	if *files != "" {
		utils.PrintInfo("Will be used for saving")
	}
	if *url != "" {
		fmt.Println("==== test ====")
		run(*url, *thread, *verbose)
	}
	if *url == "" {
		utils.PrintErr("An url should be given")
	}
}
