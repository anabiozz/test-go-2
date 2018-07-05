package bubblesort

/*
	Worst-case performance	О(n2) comparisons, О(n2) swaps
	Best-case performance	О(n) comparisons, О(1) swaps
	Average performance	О(n2) comparisons, О(n2) swaps
	Worst-case space complexity O(1) auxiliary
*/

// BubbleSort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.
func BubbleSort(a []int) []int {
	var (
		n      = len(a)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return a
}
