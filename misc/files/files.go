package files

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Os basic level
func Os() {

	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := input.Close(); err != nil {
			panic(err)
		}
	}()

	output, err := os.Create("output")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := input.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := output.Write(buf[:n]); err != nil {
			panic(err)
		}
	}
}

// Bufio package
func Bufio() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := input.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(input)

	output, err := os.Create("output")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	w := bufio.NewWriter(output)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}
}

func Ioutil() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("output", b, 0644)
	if err != nil {
		panic(err)
	}
}

func Write() {
	buf := make([]byte, 0)
	output, err := os.OpenFile("input", os.O_RDWR, os.ModeType)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			panic(err)
		}
	}()

	for i := 0; i < 1000000; i++ {
		for j := 0; j < 90; j++ {
			buf = append(buf, 0x50)
		}
		buf = append(buf, 0x0D)
	}

	err = binary.Write(output, binary.LittleEndian, buf)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
}


