package main

import (
	"fmt"
	"strconv"
	"test/interfaces/binary"
)

// Stringer ..
type Stringer interface {
	String() string
}

// ToString ..
func ToString(any interface{}) string {
	// if v, ok := any.(Stringer); ok {
	// 	return v.String()
	// }
	switch v := any.(type) {
	case Stringer:
		return v.String()
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	}
	return "???"
}

func main() {
	b := binary.Binary(200)
	s := ToString(b)
	fmt.Println(s)
}
