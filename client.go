package client

import (
	"net/http"
	"time"
)

// HTTP implements a sleepy client for HTTP requests
type HTTP struct{}

// POST some stuff to a URL
func (h *HTTP) POST(uri string, data []byte) (*http.Response, error) {

	time.Sleep(30 * time.Second)
	return *http.Response{}, nil
}
