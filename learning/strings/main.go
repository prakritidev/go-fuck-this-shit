package main

import (
	"fmt"
)

// A Go string is a read-only slice of bytes.
// The language and the standard library treat strings specially - as containers of text encoded in UTF-8. In other languages, strings are made of “characters”.
// In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point
func main() {
	const s = "fuck you"

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
