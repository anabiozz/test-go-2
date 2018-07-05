package main

import (
"crypto/md5"
"encoding/hex"
"fmt"
"log"
"runtime"
"time"
)

var hash [16]byte

func nextByte(b byte) byte {
	switch b {
	case 'z':
		return '0'
	case '9':
		return 'a'
	default:
		return b + 1
	}
}

func nextPass(b []byte) {
	for i := len(b) - 1; i >= 0; i-- {
		b[i] = nextByte(b[i])
		if b[i] != '0' {
			return
		}
	}
}

type part struct {
	start string
	end   byte
}

// worker по порядку сравнивает хэш каждого пароля с искомым
func worker(in <-chan part, out chan<- string) {
	var p part
	var b []byte
	for {
		p = <-in
		b = []byte(p.start)
		for b[0] != p.end {
			if md5.Sum(b) == hash {
				out <- string(b)
				return
			}
			nextPass(b)
		}
	}
}

func generator(in chan<- part) {
	start := []byte("00000")
	var b byte
	for {
		b = nextByte(start[0])
		in <- part{string(start), b}
		start[0] = b
	}
}

func main() {
	t := time.Now()
	const hashString = "95ebc3c7b3b9f1d2c40fec14415d3cb8" // "zzzzz"
	h, err := hex.DecodeString(hashString)
	if err != nil {
		log.Fatal(err)
	}
	copy(hash[:], h)
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	in := make(chan part)
	out := make(chan string)
	go generator(in)
	for i := 0; i < num; i++ {
		go worker(in, out)
	}
	fmt.Println("Пароль: ", <-out)
	fmt.Println("Время поиска: ", time.Since(t))
}
