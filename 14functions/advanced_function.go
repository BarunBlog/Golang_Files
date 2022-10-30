package main

import "fmt"

func test(x int) {
	fmt.Println("Hello")
}

// 		   parameter_name, fun type, func return type
func test3(myfunc func(int) int) {
	fmt.Println(myfunc(7))
}

func main() {
	x := test // assigning the reference of the test function to the x variable
	x(10)     // same as test(10)

	// saving function to a variable
	test1 := func(x int) {
		fmt.Println(x)
	}
	test1(20)

	// sending params to the function before assigning it

	test2 := func(x int) int {
		return x * -1
	} //(21)

	//fmt.Println(test2)

	test3(test2)
}
