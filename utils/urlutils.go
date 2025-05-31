package utils

import (
	"fmt"
)

const Red = "\033[31m"
const Green = "\033[32m"
const Blue = "\033[34m"
const Reset = "\033[0m"

func InitialInfo(url string, wordlist string, threads int, version string) {
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
`, version, url, wordlist, threads)
}

func StatusColor(status int) string {
	switch {
	case status >= 100 && status < 200:
		return "\033[36m" // Cyan
	case status >= 200 && status < 300:
		return "\033[32m" // Green
	case status >= 300 && status < 400:
		return "\033[34m" // Blue
	case status >= 400 && status < 500:
		return "\033[33m" // Yellow
	case status >= 500:
		return "\033[31m" // Red
	default:
		return "\033[35m" // Magenta for unknown
	}
}

func IsForbidden(statusCode int)  bool {
	switch statusCode {
	case 404:
		return true
	case 500:
		return true
	default:
		return false

	}
}
