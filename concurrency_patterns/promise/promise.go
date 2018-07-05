package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Body  []byte
	Error error
}

func promise(url string) <-chan data {
	c := make(chan data, 1)

	go func() {
		var body []byte
		var err error

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)

		c <- data{Body: body, Error: err}
	}()

	return c
}

func main() {
	future := promise("http://golang.org")

	body := <-future
	fmt.Printf("responseL %#v", string(body.Body))
	fmt.Printf("error: %#v", body.Error)
}
