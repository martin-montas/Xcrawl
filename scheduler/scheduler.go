package scheduler

import (
	"fmt"
	"net/http"
	"sort"

	"nock/worker"
)

var Links []worker.Link

func IsPathAlive(url string) (bool, int) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Domain is unreachable %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return false, response.StatusCode
	} else {
		return true, response.StatusCode
	}
}

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

func AppendToLink(l *worker.Link) {
	Links = append(Links, *l)
}
