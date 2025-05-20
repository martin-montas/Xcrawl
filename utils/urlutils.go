package utils

// func  ToFile() {}

import (
	"fmt"
	"strings"
	"time"
)

func Banner() {
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

}

func PrintErr(s string) {
	fmt.Println(
		time.Now().Format("2006-01-02 03:04:05 PM"), "[\033[32mERR\033[0m] %s", s)

}
func PrintInfo(s string) {
	fmt.Println(
		time.Now().Format("2006-01-02 03:04:05 PM"), "[\033[33mINFO\033[0m] %s", s)
}


func FullURL(s string, domain string) string {
	if strings.HasSuffix(s, "/") {

		return domain + s
	}else {

		return domain + "/" + s
	}
}



