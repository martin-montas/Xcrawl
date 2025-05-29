package fetch

import (
	"fmt"
	"net/http"
	"os"
)

func (l *Link) DisplayInfo() {
	var (
		statusCodeColored = [5]string{
			"\033[34m", // Blue
			"\033[33m", // Yellow
			"\033[32m", // Green
			"\033[31m", // Red
			"\033[0m",  // Reset
	}
		statusColor string
	)
	switch l.StatusCode {
		
		case 200:
			statusColor = statusCodeColored[2]
		case 301:
			statusColor = statusCodeColored[1]
		case 404:
			statusColor = statusCodeColored[3]
		default:
			statusColor = statusCodeColored[0]

	}
	fmt.Printf("%s[%d]%s: %s \n", statusColor, l.StatusCode, statusCodeColored[4], l.Path)
}

func CheckStatuscodeFromURL(u string)  int {
	response, err := http.Get(u)

	if err != nil {
		fmt.Printf("Domain is unreachable %s\n", u)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return  response.StatusCode
	}
	return  response.StatusCode
}

