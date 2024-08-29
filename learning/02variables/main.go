package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}

