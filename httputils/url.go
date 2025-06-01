package httputils

import (
	//"bufio"
	"fmt"
	//"io"
	//"net/http"
	"os"
	//"sync"
	//"time"
)

type URL struct {
	Name          string
	StatusCode    int
	ContentLength int64
}

func NewURL(u string) URL {
	resp, err := FetchResponse(u)
	if err != nil {
		fmt.Printf("domain is unreachable %s\n", u)
		os.Exit(1)
	}
	defer resp.Body.Close()
	size := resp.ContentLength
	if size == -1 {
		size = 3487
	}
	return URL{
		Name:          u,
		StatusCode:    resp.StatusCode,
		ContentLength: size,
	}
}
