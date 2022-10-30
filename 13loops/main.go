package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednessday", "Friday", "Saturday"}

	fmt.Println(days)

	// Unfortunately the language specification doesn't allow you to declare the variable type in the for loop.
	/*for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}*/

	/*for i := range days {
		fmt.Println(days[i])
	}*/

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day) // v – formats the value in a default format.
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Going to youtube.com")
}
