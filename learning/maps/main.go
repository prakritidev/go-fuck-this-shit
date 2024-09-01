package main

import (
	"fmt"
	"maps"
)

func plus(a, b int) int {
	return a + b
}

func main() {
	res := plus(1, 2)
	fmt.Println(" 1+2 =", res)

	fmt.Println("==============")
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("map:", m)
	v1 := m["k1"]
	fmt.Print(v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}

	n2 := map[string]int{"foo": 1, "bar": 2}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
