package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.LockOSThread()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
		fmt.Println("")
	}()

	go func() {
		defer wg.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
		fmt.Println("")
	}()

	wg.Wait()
}
