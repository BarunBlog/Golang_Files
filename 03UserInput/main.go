package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err syntax, works like try catch
	input, _ := reader.ReadString('\n') // first part is try and second is error catch
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)
}
