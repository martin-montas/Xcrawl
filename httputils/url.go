package response

import (
	//"bufio"
	"fmt"
	//"io"
	//"net/http"
	"os"
	//"sync"
	//"time"
)

type ResponseData struct {
	Name          string
	StatusCode    int
	ContentLength int64
}

func NewRequest(u string) ResponseData {
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
	return ResponseData{
		Name:          u,
		StatusCode:    resp.StatusCode,
		ContentLength: size,
	}
}
