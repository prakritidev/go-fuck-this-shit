package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}