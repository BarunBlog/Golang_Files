package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proResult, message := proAdder(2, 5, 7, 9)
	fmt.Println("Pro result is:", proResult)
	fmt.Println("Pro message is:", message)

}

func greeter() {
	fmt.Println("Namstey from golang")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi from pro function"
}
