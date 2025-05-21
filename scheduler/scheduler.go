package scheduler

import (
	"fmt"
	"net/http"
	"nock/worker"
	"sort"
)

var Links []worker.Link

func FindDuplicates() []worker.Link {
	var newLinks []worker.Link

	sort.Slice(Links, func(i, j int) bool {
		return Links[i].ID < Links[j].ID
	})
	for i, n1 := range Links {
		for j, n2 := range Links {
			if i == j {
				continue
			}
			if n1.ID == n2.ID {
				newLinks = append(newLinks, n1)
			}
		}
	}
	return newLinks
}

func IsPathAlive(url string) (bool, int) {
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

func AppendToLink(l *worker.Link) {
	Links = append(Links, *l)
}
