package main

import ( 
	"fmt"
	"time"
	"flag"
)

var urls []string
var rootDomain string

func main() {
	url 				:= flag.String("u", "", "Name to greet")
	verbose 			:= flag.Bool("v", false, "Enable verbose output.")
	thread 				:= flag.Int("t", 3, "The amount of threads.")
	files 				:= flag.String("o", "", "Output file.")
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

		coder: @github.com/martin-montas

		`)
	fmt.Print("\033[0m")
	if *verbose {
		fmt.Println(
			time.Now().Format("2006-01-02 03:04:05 PM"),
			"[\033[32mINFO\033[0m] Verbose mode on")
	} 
	if !*verbose {
		fmt.Println(
			time.Now().Format("2006-01-02 03:04:05 PM"),
			"[\033[32mINFO\033[0m] Verbose mode off")
	}

	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"),
		"[\033[32mINFO\033[0m] using",
		*thread,
		"threads")

	if *files != "" {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"),
			"[\033[32mINFO\033[0m]", *files, "will be used for saving")
	}

	if *url != "" {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"),
			"[\033[32mINFO\033[0m]", *url, "Will be used as root domain.")
		rootDomain = *url
		run(rootDomain)

		for _, val := range links {
			fmt.Println(val)
		}
	}

	if *url == "" {
		fmt.Println("An URL should be given.")
	}
}
