package main

import ( 
		"fmt"
		"log"
		"time"
		"flag"
)

var urls []string

func main() {
	url 				:= flag.String("u", "https://example.com", "Name to greet")
	verbose 			:= flag.Bool("v", true, "Enable verbose output.")
	// thread 				:= flag.Int("t", 3, "The amount of threads.")
	flag.Parse()
	fmt.Print("\033[34m")
	fmt.Println(`

		███╗   ██╗ ██████╗  ██████╗██╗  ██╗
		████╗  ██║██╔═══██╗██╔════╝██║ ██╔╝
		██╔██╗ ██║██║   ██║██║     █████╔╝ 
		██║╚██╗██║██║   ██║██║     ██╔═██╗ 
		██║ ╚████║╚██████╔╝╚██████╗██║  ██╗
		╚═╝  ╚═══╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝
		
		Nock Url/Param Crawler.
		v0.0.1

	`)
	fmt.Print("\033[0m")
	if *verbose {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM") + " [\033[32mINFO\033[0m] Verbose mode on")
	}

	urls = append(urls, *url)
	links, err := extractLinks()
	if err != nil {
		log.Fatal(err)
	}
			for _, val := range links {
				fmt.Println(val)
			}
}
