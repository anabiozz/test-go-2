package main

import (
	"math/rand"
	"testing"
)

func TestInsertSort(t *testing.T) {
	x := []int{}
	for index := 0; index < 1000000; index++ {
		x = append(x, rand.Intn(1000000))
	}

	t.Error(insertSort(x))
}

func BenchmarkInsertSort(b *testing.B) {
	x := []int{}
	for index := 0; index < 1000000; index++ {
		x = append(x, rand.Intn(1000000))
	}
	for n := 0; n < b.N; n++ {
		insertSort(x)
	}
}
