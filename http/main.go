package main

import "fmt"
import "sync"
import "io"
import "io/ioutil"
import "net"
import "net/http"
import "time"

var httpTransport http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout:   10 * time.Second,
	ResponseHeaderTimeout: 30 * time.Second,
}

// HTTPClient is a client fot http that you should re-use
var HTTPClient = http.Client{Transport: httpTransport, Timeout: 500 * time.Millisecond}

func main() {
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func() {
			for {
				request, _ := http.NewRequest(
					"GET",
					"URL",
					nil,
				)

				request.Header.Set("Accept", "application/json")
				request.Header.Set("ContentType", "application/json")

				res, err := HTTPClient.Do(request)
				if err != nil {
					fmt.Printf("ERROR: %s\n\n", err.Error())
					continue
				}
				if res.StatusCode != 204 && res.StatusCode != 200 {
					fmt.Printf("StatusCode: %d\n", res.StatusCode)
				}
				if res.Close {
					fmt.Println("connection close!")
				}
				_, err = io.Copy(ioutil.Discard, res.Body)
				if err != nil {
					fmt.Printf("ERROR: %s\n\n", err.Error())
				}
				res.Body.Close()
			}
			w.Done()
		}()
	}
	w.Wait()
}
