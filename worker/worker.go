package worker

import (
	"fmt"
	"golang.org/x/net/html"
)

type Link struct {
	parent 		html.Node

	alive      	bool
	visited    	bool
	statusCode 	int
	text       	string
}

func (l *Link) setOwner(parent html.Node, alive bool, text string, statusCode int, visited bool) {
	l.parent = parent
	l.alive = alive
	l.text = text
	l.visited = visited
	l.statusCode = statusCode
}


func (l *Link) DisplayPath() {
	fmt.Println(l.text)
}
