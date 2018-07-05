package main

import (
	"fmt"
	"log"
)

// func recursion(number int) int {
// 	if number == 1 {
// 		return number
// 	}
// 	return number + recursion(number-1)
// }

func tailRecursion(number int, product int) int {
	product = product + number
	log.Printf("number: %d\n", number)
	log.Printf("product: %d\n", product)
	if number == 1 {
		return number
	}
	return tailRecursion(number-1, product)

}

func recursiveChannel(number int, product int, result chan int) {
	log.Printf("number: %d\n", number)
	log.Printf("product: %d\n", product)
	product = product + number
	if number == 1 {
		result <- product
		return
	}
	go recursiveChannel(number-1, product, result)
}

func main() {
	// answer := recursion(4)
	// log.Println(answer)

	// answerTail := tailRecursion(4, 0)
	// log.Printf("recursive: %d\n", answerTail)

	result := make(chan int)

	recursiveChannel(4, 0, result)
	answer := <-result
	fmt.Printf("recursive: %d\n", answer)
}
