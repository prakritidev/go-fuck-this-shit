package main

import (
	"fmt"
)

/*
What is defer and when to use it ?

1. Defer is used to run when the function its called inside returns
2. Defer statement is stored in stack.
3. This is generalyy used for cleanup actions if necessary
	examples: Database connections cleanup or waitgorup cleanup etc.



*/


func a(){
	i := 0
	defer fmt.Print(i)
	i++
	fmt.Println("\n")
	return

}


func b() {
	
	for i:= 0; i < 4; i++ {
		defer fmt.Print(i)
	}
	
}

func c() {
	defer fmt.Println("Yohoho")
	defer fmt.Println("Yohoho1")
	defer fmt.Println("Yohoho2")
	fmt.Println("duck")
	fmt.Println("\n")

}

func d() (i int) {
    defer func() { i++ }()
    return 1
}


func main(){
//  Exmpale 1:
	// defer a()
// Example 2: A deferred function’s arguments are evaluated when the defer statement is evaluated.
	// defer b()
// Example 3: Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	// defer c()
// Example 4: Deferred functions may read and assign to the returning function’s named return values.
	defer fmt.Println(d())

}