package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var called = false

func setupTests() (*httptest.Server, func()) {
	called = false

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("ok"))
			},
		),
	)

	return ts, func() {
		ts.Close()
	}
}

func TestGet(t *testing.T) {
	ts, done := setupTests()
	defer done()

	sc := HTTP{}
	resp, err := sc.GET(ts.URL)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Expected status OK, got:", resp.StatusCode)
	}
}
