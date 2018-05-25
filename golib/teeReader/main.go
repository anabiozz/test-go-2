package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Reader ...
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer ..
type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {

	// resp, err := http.Get("http://www.golangprograms.com/example-readall-readdir-and-readfile-from-io-package.html")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// reader, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	fo, err := os.Create("output")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(fo)

	testString := strings.NewReader("Jobs, Code, Videos and News for Go hackers.")
	example := io.TeeReader(testString, w)
	readerMap := make([]byte, testString.Len())
	fmt.Println(readerMap)
	length, err := example.Read(readerMap)
	fmt.Println(length)
}
