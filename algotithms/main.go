package main

import (
	"log"
	"test/algotithms/sort/bubblesort"
)

func main() {
	// slice := utils.GenerateSlice(10)
	x := []int{1, 83, 101, 43, 4, 5}
	// log.Println(quicksort.QuickSort(slice))
	log.Println(bubblesort.BubbleSort(x))
}