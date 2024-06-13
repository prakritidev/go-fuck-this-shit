package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("DuckMe!!")
	fmt.Println(message)
}
