package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprint("Hi, %v Welcome!", name)
	return message
}
