package main

import (
	"cmp"
	"fmt"
	"slices"
)

// These are like java comaparator.
func main() {

	fruits := []string{"peach", "banana"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)

	fmt.Println(fruits)
}
