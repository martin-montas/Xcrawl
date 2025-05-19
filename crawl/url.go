package crawl

import "fmt"

type Link struct {
	path   string
	lives  bool
	status int
	text 	 string
}

func (u *Link) SetOwner(path string, lives bool) {
	u.path = path
	u.lives = lives

}
func (u *Link) Owner() Link {
	return *u
}

func (u *Link) PrintLink() {
	fmt.Println(u.path)

}
func (u *Link) Printlives() bool {
	fmt.Println(u.lives)
	return true
}
