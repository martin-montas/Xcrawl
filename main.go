package main

import ( 
		"fmt"
		"time"
		"flag"
)

func main() {
	url 				:= flag.String("u", "https://example.com", "Name to greet")
	verbose 			:= flag.Bool("v", true, "Enable verbose output.")
	thread 				:= flag.Int("t", 3, "The amount of threads.")
	flag.Parse()

	if *verbose {
		fmt.Println("[" + time.Now().Format("2006-01-02 03:04:05 PM") + "]" +  " [INFO] Verbose mode on")
	}

	fmt.Println("Hello,", *url)
	fmt.Println("Hello,", *thread)
}
