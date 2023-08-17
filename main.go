package main

import "fmt"

func main(){
	fmt.Println("Yohohohoho!!! watashi wa brook desu");
	sever := server(":8080")
	sever.Run()
}