package insertsort

/*
	1) Select the first unsorted element
	2) Swap other elements to the right to create the correct position and shift the unsorted element
	3) Advance the marker to the right one element

	Worst-case performance	О(n2) comparisons, swaps
	Best-case performance	O(n) comparisons, O(1) swaps
	Average performance	О(n2) comparisons, swaps
	Worst-case space complexity О(n) total, O(1) auxiliary
*/

// InsertionSort ...
func InsertionSort(a []int) []int {
	var length = len(a)
	for i := 1; i < length; i++ {
		marker := i
		for marker > 0 {
			if a[marker-1] > a[marker] {
				a[marker-1], a[marker] = a[marker], a[marker-1]
			}
			marker = marker - 1
		}
	}
	return a
}
