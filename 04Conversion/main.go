// https://www.linkedin.com/pulse/mastering-golang-series-part-1-how-read-data-from-file-wang-ph-d-/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Taking string as an input and converting it as an integer

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // TrimSpace function to remove leading and trailing whitespace as defined by Unicode.

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

	text_reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text here: ")

	line, _, err := text_reader.ReadLine() // enter "Emma"

	if err != nil { // check errorpanic(err)
	}

	fmt.Println(line) // prints [69 109 109 97]

	fmt.Println(string(line)) // prints Emma
}
