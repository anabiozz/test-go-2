package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromise(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plane")
		io.WriteString(w, "Golang Concurrancy")
	}))
	defer ts.Close()

	future := promise(ts.URL)
	expectedChan := data{Body: []byte("Golang Concurrancy"), Error: nil}

	assert.Equal(t, expectedChan, <-future)
}
