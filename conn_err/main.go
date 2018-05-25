package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func startWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8080", nil)
}

func startLoadTest() {
	count := 0
	for {
		resp, err := http.Get("http://www.philipotoole.com")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		resp.Body.Close()

		rebo, _ := ioutil.ReadAll(resp.Body) // if we do not read from body it does not close connection and create many TIME_WAIT сonnection

		log.Println(rebo)
		log.Printf("Finished GET request %#v", count)
		count++
		time.Sleep(time.Second * 1)
	}
}

func main() {
	startWebServer()
	startLoadTest()
}
