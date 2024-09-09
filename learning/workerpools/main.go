package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done \n", id)
}

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				ops.Add(1)
			}
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
