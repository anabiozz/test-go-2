package main

import (
	"context"
	"sync"
)

type TaskFunction func() interface{}

func PerformTasks(ctx context.Context, tasks []TaskFunction) chan interface{} {

	workers := make([]chan interface{}, 0, len(tasks))

	for _, task := range tasks {
		resultChannel := newWorker(ctx, task)
		workers = append(workers, resultChannel)
	}

	out := merge(ctx, workers)
	return out
}

func newWorker(ctx context.Context, task TaskFunction) chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		select {
		case <-ctx.Done():
			return
		case out <- task():
		}
	}()
	return out
}

func merge(ctx context.Context, workers []chan interface{}) chan interface{} {
	out := make(chan interface{})

	var wg sync.WaitGroup

	output := func(c <-chan interface{}) {
		defer wg.Done()
		for result := range c {
			select {
			case <-ctx.Done():
				return
			case out <- result:
			}
		}
	}

	wg.Add(len(workers))

	for _, workerChannel := range workers {
		go output(workerChannel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
