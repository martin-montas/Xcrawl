package nockcrawl

import (
	"flag"
	"fmt"
	"log"
	"os"

	"nock/httputils"
)

type OptionsCrawl struct {
	Wordlist string
	BaseURL  string
	Threads  int
	Version  string
}

func (c *NockCrawl) Parse(version string) {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'dir' , 'crawl' or 'version' subcommand")
		return
	}
	dirCmd := flag.NewFlagSet("dir", flag.ExitOnError)
	u := dirCmd.String("u", "", "Target URL")
	t := dirCmd.Int("t", 10, "Number of threads")

	if err := dirCmd.Parse(os.Args[2:]); err != nil {
		log.Fatalf("failed to parse dir command: %v", err)
	}

	if *u == "" {
		fmt.Println("Usage: dir -u <url> -w <wordlist> -t <threads>")
		return
	}
	o := &OptionsCrawl{
		BaseURL: *u,
		Threads: *t,
		Version: version,
	}
	c.opt = o
	c.client = httputils.NewHTTPClient()
}

func (o *OptionsCrawl) DisplayBanner() {
	fmt.Printf(`
===============================================================
xcrawl %-6s 
by martin montas - @github.com/martin-montas
===============================================================
[+] URL:      		%-21s
[+] Wordlist: 		%-21s
[+] Threads: 		%-21d
===============================================================
                       STARTING                       
===============================================================
`, o.Version, o.BaseURL, o.Wordlist, o.Threads)
}
