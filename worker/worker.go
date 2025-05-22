package worker

import (
	"fmt"
)

type Link struct {
	Alive      bool
	StatusCode int
	Path       string
	ID         int
}

var Links []Link


func (l *Link) DisplayPath() {
	fmt.Println(l.Path)
}
