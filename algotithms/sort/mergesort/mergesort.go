package mergesort

/*
	Merge Sort is a Divide and Conquer algorithm. It divides input array in two halves,
	calls itself for the two halves and then merges the two sorted halves.
	The merge() function is used for merging two halves. The merge(arr, l, m, r) is key process that assumes that arr[l..m]
	and arr[m+1..r] are sorted and merges the two sorted sub-arrays into one. See following C implementation for details.

	Worst-case performance	O(n log n)
	Best-case performance O(n log n) typical, O(n) natural variant
 	Average performance	O(n log n)
 	Worst-case space complexity	О(n) total with O(n) auxiliary, O(1) auxiliary with linked lists[1]
*/

// MergeSort ...
func MergeSort(a []int) []int {
	n := len(a)

	if n == 1 {
		return a
	}

	middle := int(n / 2)

	var (
		left  = make([]int, middle)
		right = make([]int, n-middle)
	)

	for i := 0; i < n; i++ {
		if i < middle {
			left[i] = a[i]
		} else {
			right[i-middle] = a[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}
