package main

import "log"

func main() {
	n := 21
	var result uint64
	result = 1

	for i := 1; i <= n; i++ {
		result *= uint64(i)
	}

	log.Println(result)
}
