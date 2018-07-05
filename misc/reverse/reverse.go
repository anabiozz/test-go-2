package main

import (
	"fmt"
	"math"
)

func reverse(x []int) []int {
	n := len(x)
	for i := 0; i <= int(math.Floor(float64((n-2)/2))); i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
	return x
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverse(x))
	fmt.Println(reverseString("Hello World"))
}
