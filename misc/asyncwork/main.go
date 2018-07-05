package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	slowFunction := func() interface{} {
		time.Sleep(time.Second * 10)
		fmt.Println("slow function")
		return 2
	}

	verySlowFunction := func() interface{} {
		time.Sleep(time.Second * 5)
		fmt.Println("very slow function")
		return "I'm ready"
	}

	errorFunction := func() interface{} {
		time.Sleep(time.Second * 1)
		fmt.Println("function with an error")
		return errors.New("Error in function")
	}

	tasks := []TaskFunction{slowFunction, verySlowFunction, errorFunction}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultChannel := PerformTasks(ctx, tasks)

	for result := range resultChannel {
		switch result.(type) {
		case error:
			fmt.Println("Received error")
			// cancel()
			// return
		case string:
			fmt.Println("Here is a string: ", result.(string))
		case int:
			fmt.Println("Here is a interger: ", result.(int))
		default:
			fmt.Println("Some unknown type: ")
		}
	}
}
