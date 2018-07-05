package main

import (
	"log"
)

// IterFib Fastest
func IterFib(n uint64) (result uint64) {
	for i, first, second := 0, 0, 1; uint64(i) <= n; i, first, second = i+1, first+second, first {
		if uint64(i) == n {
			result = uint64(first)
		}
	}
	return result
}

// RecurFib Slowest
func RecurFib(n int) int {
	if n < 2 {
		return n
	}
	return RecurFib(n-2) + RecurFib(n-1)
}

func notRecursiveFib(n int) int {
	if n == 0 {
		return 1
	}
	curr := 1
	prev := 1
	for i := 2; i <= n; i++ {
		tmp := curr
		curr += prev
		prev = tmp
	}
	return curr
}

// TailRecurFib so-so
func TailRecurFib(n int) int {
	return fibo(n, 1)
}

func fibo(n, result int) int {
	if n == 0 {
		return 1
	}
	return fibo(n-1, n*result)
}

// ConcurrencyFibo ...
func ConcurrencyFibo(n int) chan int {
	return concurrencyFibo(n)
}

func concurrencyFibo(n int) chan int {
	result := make(chan int, 0)
	go func() {
		defer close(result)
		if n <= 2 {
			result <- 1
			return
		}
		result <- <-concurrencyFibo(n-1) + <-concurrencyFibo(n-2)
	}()
	return result
}

func main() {
	log.Println(IterFib(5))
	// for i := range ConcurrencyFibo(6) {
	// 	fmt.Printf("%d\n", i)
	// }
}
