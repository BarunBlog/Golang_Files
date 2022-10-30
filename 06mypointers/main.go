package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	//var ptr *int

	//fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber // var ptr *int = &myNumber same
	fmt.Println("address of myNumber is", ptr)
	fmt.Println("Value of the pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)
}
