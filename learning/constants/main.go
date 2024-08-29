package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 50000000000

	fmt.Print(math.Sin(n))
	for n := range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println("Duck %d", n)
	}
}
