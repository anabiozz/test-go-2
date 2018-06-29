package main

import "testing"

func TestReverseInPlace(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6}
	res := ReverseInPlace(x)
	t.Error(res)
}
