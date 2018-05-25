package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.wildberries.ru", nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info %+v\n", dnsInfo)
		},
		WroteRequest: func(wroteRequest httptrace.WroteRequestInfo) {
			fmt.Printf("WroteRequest Info %+v\n", wroteRequest)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
