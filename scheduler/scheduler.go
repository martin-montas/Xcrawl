package scheduler

import (
	"fmt"
	"net/http"
	// "nock/worker"
)

func IsLinkAlive(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Domain is unreachable %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	} else {
		return true
	}
}

func IsDuplicate() bool {
	// for now...
	return true
}
