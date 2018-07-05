package bubblesort

import (
	"log"
	"test/algotithms/utils"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(utils.GenerateSlice(1000))
	}
}

func TestMergeSort(t *testing.T) {
	sorted := BubbleSort(utils.GenerateSlice(1000))
	log.Println(sorted)
	if len(sorted) != 1000 {
		t.Error("wrong")
	}
}
