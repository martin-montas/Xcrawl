package main

import (
	"flag"
	"nock/Worker"
	"nock/features"
)

func main() {
	u := flag.String("u", "", "Name to greet")
	v := flag.Bool("v", false, "Enable verbose output.")
	t := flag.Int("t", 3, "The amount of threads.")
	files := flag.String("o", "", "Output file.")
	flag.Parse()
	features.Banner()
	if *v {
		features.PrintInfo("Verbose mode is on")
	}
	if !*v {
		features.PrintInfo("Verbose mode is off")
	}
	features.PrintInfo("Using threads given")
	if *files != "" {
		features.PrintInfo("Will be used for saving")
	}
	if *u != "" {
		worker.Run(*u, *t, *v)
	}
	if *u == "" {
		features.PrintErr("An url should be given")
	}
}
