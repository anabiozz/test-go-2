package reverse

import (
	"math"
)

func ReverseInPlace(x []int) []int {
	n := len(x)
	for i := 0; i <= int(math.Floor(float64((n-2)/2))); i++ {
		tmp := x[i]
		x = append(x[:i], x[n-1-i])
		x = append(x[:n-1-i], tmp)
	}
	return x
}
