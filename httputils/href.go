package response

import (
	"fmt"
	"net/http"
	"os"
)

type Href struct {
	URL            string
	ResponseLength int64
	response       *http.Response
	url            *ResponseData
}

func NewHref(url string) Href {
	fetchResponse, err := FetchResponse(url)
	if err != nil {
		fmt.Printf("domain is unreachable %s\n", url)
		os.Exit(1)
	}
	u := NewRequest(url)
	return Href{
		URL:            url,
		ResponseLength: fetchResponse.ContentLength,
		response:       fetchResponse,
		url:            &u,
	}
}

func FetchResponse(url string) (*http.Response, error) {
	if len(url) > 0 && url[len(url)-1] != '/' {
		url += "/"
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("domain is unreachable: %s", url)
	}
	return resp, nil
}
