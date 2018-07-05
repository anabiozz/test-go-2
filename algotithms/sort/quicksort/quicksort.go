package quicksort

/*
	1) Pick an element, called a pivot, from the array.
	2) Partitioning: reorder the array so that all elements with values less than the pivot come before the pivot,
		while all elements with values greater than the pivot come after it (equal values can go either way).
		After this partitioning, the pivot is in its final position. This is called the partition operation.
	3) Recursively apply the above steps to the sub-array of elements with smaller values and separately to the sub-array of elements with greater values.

	Worst-case performance	O(n2)
	Best-case performance	O(n log n) (simple partition) or O(n) (three-way partition and equal keys)
	Average performance	O(n log n)
	Worst-case space complexity	O(n) auxiliary (naive) O(log n) auxiliary (Sedgewick 1978)
*/

// QuickSort ...
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}
