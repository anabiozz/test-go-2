package main

import "testing"

func BenchmarkReverse(b *testing.B) {
	x := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		reverse(x)
	}
}

func BenchmarkReverseString(b *testing.B) {
	s := "Hello World"
	for i := 0; i < b.N; i++ {
		reverseString(s)
	}
}
