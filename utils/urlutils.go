package utils

// TODO()
// func ToFile() {}

import (
	"fmt"
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

func PrintErr(s , value string) {
	fmt.Printf(
		time.Now().Format("2006-01-02 03:04:05 PM"), "[\033[31mERR\033[0m] %s \n", value)

}
func PrintInfo(s , value string) {
	fmt.Printf(
		time.Now().Format("2006-01-02 03:04:05 PM"), "[\033[33mINFO\033[0m] %s\n", value)
}




