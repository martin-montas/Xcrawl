package nockdir

import (
	"fmt"
)

const defaultList = "/usr/share/dirb/wordlists/common.txt"

type OptionsDir struct {
	Wordlist string
	BaseURL  string
	Threads  int
}

// displays the initial information
func (o *OptionsDir) DisplayBanner() {
	fmt.Printf("Wordlist: %s\n", o.Wordlist)
	fmt.Printf("Base URL: %s\n", o.BaseURL)
	fmt.Printf("Threads: %d\n", o.Threads)
}
