package worker

import (
	"fmt"
	"golang.org/x/net/html"
)

type Link struct {
	Alive      bool
	StatusCode int
	Path       string
	ID         int

	Node 			html.Node
}

var Links []Link


func (l *Link) DisplayPath() {
	fmt.Println(l.Path)
}

