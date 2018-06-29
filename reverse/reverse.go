package main

import (
	"math"
)

// 	x := []int{1, 2, 3, 4, 5, 6}
func ReverseInPlace(x []int) []int {
	n := len(x)
	for i := 0; i <= int(math.Floor(float64((n-2)/2))); i++ {
		tmp := x[i]
		x = append(x[:i], x[n-1-i])
		x = append(x[:n-1-i], tmp)
	}
	return x
}
