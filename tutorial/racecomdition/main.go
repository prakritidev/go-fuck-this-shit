package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("RaceCondition Stuff")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	
	var score = []int{0}

	wg.Add(3)

	go func (wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	} (wg, mut)

	go func (wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two T")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()	
	} (wg, mut)
		

	go func (wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three F")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	} (wg, mut)

	wg.Wait()
	fmt.Println(score)
}