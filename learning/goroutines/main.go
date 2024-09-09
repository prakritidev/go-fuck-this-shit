package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func worker(done chan bool) {
	fmt.Println("wprking ....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	// something is wrong here, not sure what
	done := make(chan bool, 1)
	go worker(chan bool)

	f("direct")
	go f("goroutinex")

	go func(msg string) {
		fmt.Print(msg)
	}("Hello")

	message := make(chan string)
	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

	time.Sleep(time.Second)
	fmt.Println("done")
}
