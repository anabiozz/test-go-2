package main

import (
	"fmt"
)

func getPrintMessage() func(string) {
	return func(message string) {
		fmt.Println(message)
	}
}

func outer(name string) {
	text := "Modified " + name
	foo := func() {
		fmt.Println(text)
	}
	foo()
}

func sum(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return
}

func main() {

	// anonymous function
	printfunc := getPrintMessage()
	printfunc("lollol")

	// anonymous function
	add := func(m int) int {
		return m + 1
	}
	result := add(6)
	fmt.Println(result)

	// closure
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	addN(5)
	result2 := addN(6)
	fmt.Println(result2)

	// variadic
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
