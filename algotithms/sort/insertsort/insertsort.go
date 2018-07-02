package insertsort

/*
	1) Select the first unsorted element
	2) Swap other elements to the right to create the correct position and shift the unsorted element
	3) Advance the marker to the right one element
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
