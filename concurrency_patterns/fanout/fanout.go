package main

import (
	"fmt"
	"time"
)

func producer(iters int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= iters; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		defer close(c)
	}()
	return c
}

func consumer(cin <-chan int) {
	for i := range cin {
		fmt.Println(i)
	}
}

func fanOut(ch <-chan int, size, lag int) []chan int {
	cs := make([]chan int, size)
	for i := range cs {
		cs[i] = make(chan int, lag)
	}
	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}
		for _, c := range cs {
			close(c)
		}
	}()
	return cs
}

func fanOutUnbufered(ch <-chan int, size int) []chan int {
	cs := make([]chan int, size)
	for i := range cs {
		cs[i] = make(chan int)
	}
	go func() {
		for i := range ch {
			for _, c := range cs {
				c <- i
			}
		}
		for _, c := range cs {
			close(c)
		}
	}()
	return cs
}

func main() {
	c := producer(10)
	chans := fanOut(c, 3, 3)
	go consumer(chans[0])
	go consumer(chans[1])
	consumer(chans[2])
}
