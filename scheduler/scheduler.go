package scheduler

import (
	"fmt"
	"net/http"
	"nock/worker"
)

var links []worker.Link

func IsLinkAlive(url string) (bool, int) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Domain is unreachable %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, resp.StatusCode
	} else {
		return true, resp.StatusCode
	}
}

func appendToLink(l *worker.Link) {
	links = append(links, *l)

}

func IsDuplicate() bool {
	// checks for duplicates in the worker.Link slice
	// for now...
	return true
}
