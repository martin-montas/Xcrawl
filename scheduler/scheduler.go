package scheduler

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"

)

type Link struct {
	Alive      	bool
	StatusCode 	int
	Path       	string
	ID         	int
}

func (l *Link) DisplayInfo() {
	var statusCodeColored = [5]string {
		"\033[34m", 		// Blue
		"\033[33m", 		// Yellow
		"\033[32m", 		// Green
		"\033[31m", 		// Red
		"\033[0m",			// Reset
	}

	var statusColor string
	if l.StatusCode <= 100 && l.StatusCode <= 101 {
		statusColor = statusCodeColored[0]

	} else if l.StatusCode <= 200 && l.StatusCode <= 204 {
		statusColor = statusCodeColored[2]

	} else if l.StatusCode <= 301 && l.StatusCode <= 304 {
		statusColor = statusCodeColored[1]

	} else {
		statusColor = statusCodeColored[3]
	}
	fmt.Printf("%s[%d]%s : %s \n",statusColor,l.StatusCode, statusCodeColored[4], l.Path)
}

var Nodes 	[]html.Node
var Links 	[]Link

func IsPathAlive(url string) (bool, int) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Domain is unreachable %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return false, response.StatusCode
	} 
	return true, response.StatusCode
}

func AppendToNode(n *html.Node) {
	Nodes = append(Nodes, *n)
}

func AppendToLinks(l *Link) {
	Links = append(Links, *l)
}

func ReturnNodes() []html.Node {
	return Nodes
}
