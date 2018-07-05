package main

import (
	"fmt"
	"strconv"
)

func print() {
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func main() {
	s := "hello, world"

	//  Atoi converts a string to an int
	a, err := strconv.Atoi("213")
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi("412")
	if err != nil {
		fmt.Println(err)
	}
	result := sum(a, b)

	print()
	fmt.Println(s[:4])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(result)
}

func sum(a, b int) int {
	return a + b
}
