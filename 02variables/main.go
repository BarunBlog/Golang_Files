package main

import "fmt"

const LoginToken string = "skdjfbsdfs5df" // Public

func main() {
	var username string = "Barun"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float32 = 255.455591155645464
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var largeFloat float64 = 255.455591155645464
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T\n", largeFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable is of type: %T\n", anotherStringVariable)

	// implicit type
	var website = "youtube.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
