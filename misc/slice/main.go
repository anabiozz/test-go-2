package main

import "log"

func main() {
	i := 0
	j := 4

	integer := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	type s struct{}

	obj := s{}
	pointerToInteger := []*s{&obj, &obj, &obj, &obj, &obj, &obj, &obj, &obj, &obj, &obj}

	// cut
	integer = append(integer[:i], integer[j:]...)

	// delete
	integer = append(integer[:i], integer[i+1:]...)

	// delete without preserving order
	integer[i] = integer[len(integer)-1]
	integer = integer[:len(integer)-1]

	// If the type of the element is a pointer or a struct with pointer fields, which need to be garbage collected

	// cut
	copy(pointerToInteger[i:], pointerToInteger[j:])
	for k, n := len(pointerToInteger)-j+i, len(pointerToInteger); k < n; k++ {
		pointerToInteger[k] = nil
	}
	pointerToInteger = pointerToInteger[:len(pointerToInteger)-j+i]

	// delete
	copy(pointerToInteger[i:], pointerToInteger[i+1:])
	pointerToInteger[len(pointerToInteger)-1] = nil
	pointerToInteger = pointerToInteger[:len(pointerToInteger)-1]

	// delete without preserving order
	pointerToInteger[i] = pointerToInteger[len(pointerToInteger)-1]
	pointerToInteger[len(pointerToInteger)-1] = nil
	pointerToInteger = pointerToInteger[:len(pointerToInteger)-1]

	log.Println(pointerToInteger)
}
