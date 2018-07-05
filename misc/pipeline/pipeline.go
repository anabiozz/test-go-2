package main

import "fmt"

func pipeline2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

func pipeline3() {
	counter := func(out chan<- int) {
		for x := 0; x < 100; x++ {
			out <- x
		}
		close(out)
	}

	squarer := func(out chan<- int, in <-chan int) {
		for v := range in {
			out <- v * v
		}
		close(out)
	}

	printer := func(in <-chan int) {
		for v := range in {
			fmt.Println(v)
		}
	}

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func main() {
	pipeline3()
}
