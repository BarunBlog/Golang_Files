package main

import "fmt"

func changeValue(str *string) {
	*str = "hello world!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	/*x := 7
	y := &x
	fmt.Println(x, y)

	*y = 8
	fmt.Println(x, y)*/

	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	changeValue2(toChange)
	fmt.Println(toChange)

	var pointer *string = &toChange
	fmt.Println(*pointer)
	fmt.Println(pointer, &pointer) // pointer to the toChange and pointer to the pointer pointing to the toChange
}
