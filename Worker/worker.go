package worker

import (
	"nock/Fetcher"
	"nock/Parser"
	"fmt"
)


type Link struct {
	path    string
	lives   bool
	visited bool
	status  int
	text    string
}

func Run(d string, t int, v string) {
	n := fetcher.Parse(d)
	parser.extract(*n, d)
}

func (u *Link) setOwner(path string, lives bool, text string, visited bool) {
	u.path = path
	u.lives = lives
	u.text = text
	visited = visited
}

func (u *Link) owner() Link {
	return *u
}

func (u *Link) printLink() {
	fmt.Println(u.path)

}
func (u *Link) printlives() bool {
	fmt.Println(u.lives)
	return true
}
