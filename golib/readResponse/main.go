package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
	"net"
	"net/http"
)

func IO(sock net.Conn, reqbuf *bytes.Buffer) (*http.Response, error) {
	_, e := io.Copy(sock, reqbuf)
	if e != nil {
		return nil, e
	}
	bread := bufio.NewReader(sock)
	res, err := http.ReadResponse(bread, &dummyRequest)
	if err != nil {
		return nil, err
	}
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		res.Header.Del("Content-Length")
		zr, err := gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		res.Body = gzreadCloser{zr, res.Body}
	}
	return res, nil
}

type gzreadCloser struct {
	*gzip.Reader
	io.Closer
}

func (gz gzreadCloser) Close() error {
	return gz.Closer.Close()
}

var dummyRequest http.Request

func main() {}
