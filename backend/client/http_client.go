package client

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 10_000
	t.MaxIdleConnsPerHost = 25
	t.ForceAttemptHTTP2 = true

	return &http.Client{
		Timeout:   time.Second * 3,
		Transport: t,
	}
}
