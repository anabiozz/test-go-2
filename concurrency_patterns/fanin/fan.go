package main

import (
	"fmt"
)

func main() {
	randomNumbers := []int{13, 45, 97, 65, 4, 65, 653, 2645, 6656}

	inputChan := generatePipeline(randomNumbers)

	c1 := squareNumber(inputChan)
	c2 := squareNumber(inputChan)

	c := fanIn(c1, c2)
	sum := 0

	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Printf("Total sum of square: %d\n", sum)
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

	return out
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func fanIn(input1, input2 <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
