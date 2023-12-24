package main

import (
	"fmt"
	"sync"
)

func main(){

	fmt.Println("Channels Intro")

	mych := make(chan int, 5)
	wg := &sync.WaitGroup{}


	// mych <- 5

	// fmt.Println(<- mych)
wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup)  {
		val, isChanOpen := <-mych

		fmt.Print(isChanOpen)
		fmt.Print(val)
		wg.Done()
	}(mych, wg)

	go func(ch chan<- int, wg *sync.WaitGroup)  {
		
		mych<-5
		mych<-6

		close(mych)
		
		wg.Done()
	}(mych, wg)

	wg.Wait()
}