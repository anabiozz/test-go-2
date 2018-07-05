package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 6
	expect := 11
	res := sum(a, b)
	if res != expect {
		t.Error("error")
	}
}
