package main

import (
	"fmt"
	"sync"
)

type ryanWorker struct {
	in   chan int
	done func()
}

func doWork(id int, w ryanWorker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %d \n", id, n)
		w.done()
	}
}

func createWorker2(id int, wg *sync.WaitGroup) ryanWorker {
	w := ryanWorker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo2() {
	var wg sync.WaitGroup

	var workers [10]ryanWorker
	for i := 0; i < len(workers); i++ {
		workers[i] = createWorker2(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	chanDemo2()
}
