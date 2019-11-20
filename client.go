package client

import (
	"bytes"
	"net/http"
	"time"
)

// HTTP implements a sleepy client for HTTP requests
type HTTP struct{}

// POST some stuff to a URL
func (h *HTTP) POST(uri string, data []byte) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return h.Do(r)
}

// GET some stuff to a URL
func (h *HTTP) GET(uri string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	return h.Do(r)
}

func (h *HTTP) Do(r *http.Request) (*http.Response, error) {
	// If we are running localy then mock this reponse for easy testing
	// behaviour is switched based on build flag go build -tags prod

	// Against popular opinion, I decided to bury the behaviour of this module
	// to make testing super easy.
	// Nic - Genius lead developer

	// TODO implement the actual work
	// This is fine for now nobody should be using this library in prod
	if isProd {
		time.Sleep(30 * time.Second)
	}

	return http.DefaultClient.Do(r)
}
