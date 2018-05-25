package main

import "fmt"

const size = 64

type bits uint64

type BitSet []bits

func (s *BitSet) Set(i uint) {
	if len(*s) < int(i/size+1) {
		r := make([]bits, i/size+1)
		copy(r, *s)
		*s = r
	}
	(*s)[i/size] |= 1 << (i % size)
}

func (s *BitSet) Clear(i uint) {
	if len(*s) >= int(i/size+1) {
		(*s)[i/size] &^= 1 << (i % size)
	}
}

func (s *BitSet) IsSet(i uint) bool {
	return (*s)[i/size]&(1<<(i%size)) != 0
}

func main() {
	s := new(BitSet)
	s.Set(13)
	s.Set(45)
	fmt.Printf("s.IsSet(13) = %t; s.IsSet(45) = %t; s.IsSet(30) = %t\n", s.IsSet(13), s.IsSet(45), s.IsSet(30))
}
