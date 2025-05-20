package worker

import (
	"fmt"
)

type Link struct {
	Alive      bool
	Visited    bool
	StatusCode int
	Path       string
}


func (l *Link) DisplayPath() {
	fmt.Println(l.Path)
}
