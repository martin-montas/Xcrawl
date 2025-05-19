package crawl

import "fmt"

type Link struct {
	path   string
	lives  bool
	status int
	text 	 string
}

func (u *Link) setOwner(path string, lives bool, text string) {
	u.path = path
	u.lives = lives
	u.text = text

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
