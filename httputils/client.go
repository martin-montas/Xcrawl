package httputils

import (
	"net"
	"net/http"
	"time"
)

type HTTPClient struct {
	client    *http.Client
	userAgent string
}

func NewHttpClient(userAgent string) *HTTPClient {
	client := &http.Client{
		Timeout: 10 * time.Second, // total request timeout
		Transport: &http.Transport{
			MaxIdleConns:        100,              // total idle connections
			MaxIdleConnsPerHost: 100,              // idle per-host (boosts reuse)
			MaxConnsPerHost:     100,              // hard cap of connections per host
			IdleConnTimeout:     90 * time.Second, // how long to keep idle conns alive
			DisableKeepAlives:   false,            // keep connection open (reused)
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return &HTTPClient{
		client:    client,
		userAgent: userAgent,
	}
}
