package main

import (
	"fmt"
	"testing"
)

func BenchmarkIterFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IterFib(5)
	}
}

func BenchmarkRecurFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RecurFib(5)
	}
}

func BenchmarkNotRecursiveFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		notRecursiveFib(5)
	}
}

func BenchmarkConcurrencyFibo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcurrencyFibo(5)
	}
}

func BenchmarkTailRecurFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TailRecurFib(5)
	}
}

func TestTailRecurFib(t *testing.T) {
	fmt.Println(TailRecurFib(5))
}
