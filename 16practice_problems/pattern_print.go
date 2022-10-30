package main

import "fmt"

func staircase(n int32) {
	var i int32
	var j int32
	var space int32 = n - 1

	for i = 1; i <= n; i++ {

		for j = 1; j <= space; j++ {
			fmt.Printf(" ")
		}
		for j <= n {
			fmt.Printf("#")
			j++
		}
		fmt.Println()
		space--
	}

}

func main() {
	var n int32 = 6

	staircase(n)
}
