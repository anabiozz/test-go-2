package main

func insertSort(A []int) []int {

	for j, val := range A {
		i := j - 1
		for i >= 0 && A[i] > val {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = val
	}

	return A
}
