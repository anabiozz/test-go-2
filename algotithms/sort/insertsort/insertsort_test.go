package insertsort

import (
	"test/algotithms/utils"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InsertionSort(utils.GenerateSlice(1000))
	}
}

func TestMergeSort(t *testing.T) {
	sorted := InsertionSort(utils.GenerateSlice(1000))
	if len(sorted) != 1000 {
		t.Error("wrong")
	}
}
