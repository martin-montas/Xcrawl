package worker

import (
	"fmt"
)

type Link struct {
	Alive      	bool
	Visited    	bool
	StatusCode 	int
	Path 			string
}

func (l *Link) setOwner(alive bool, path string, statusCode int, visited bool) {
	l.Alive = alive
	l.Path = path
	l.Visited = visited
	l.StatusCode = statusCode
}

func (l *Link) DisplayPath() {
	fmt.Println(l.Path)
}
