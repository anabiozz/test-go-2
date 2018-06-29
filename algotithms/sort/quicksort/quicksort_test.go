package quicksort

import (
	"test/algotithms/utils"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	slice := utils.GenerateSlice(999999)
	for n := 0; n < b.N; n++ {
		QuickSort(slice)
	}
}
