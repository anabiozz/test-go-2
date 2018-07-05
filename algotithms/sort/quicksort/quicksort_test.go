package quicksort

import (
	"test/algotithms/utils"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		QuickSort(utils.GenerateSlice(1000))
	}
}

func TestQuickSort(t *testing.T) {
	sorted := QuickSort(utils.GenerateSlice(1000))
	if len(sorted) != 1000 {
		t.Error("wrong")
	}
}
