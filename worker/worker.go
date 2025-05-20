package worker

import (
	"fmt"
)

type Link struct {
	alive      bool
	visited    bool
	statusCode int

	path string
}

func (l *Link) setOwner(alive bool, path string, statusCode int, visited bool) {
	l.alive = alive
	l.path = path
	l.visited = visited
	l.statusCode = statusCode
}

func (l *Link) DisplayPath() {
	fmt.Println(l.path)
}
