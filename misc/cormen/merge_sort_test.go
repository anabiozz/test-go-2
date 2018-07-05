package main

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	x := []int{}
	for index := 0; index < 10000000; index++ {
		x = append(x, rand.Intn(10000000))
	}
	res := MergeSort(x)
	t.Error(res)
}

func BenchmarkMergeSort(b *testing.B) {
	x := []int{}
	for index := 0; index < 1000000; index++ {
		x = append(x, rand.Intn(1000000))
	}
	for n := 0; n < b.N; n++ {
		MergeSort(x)
	}
}
