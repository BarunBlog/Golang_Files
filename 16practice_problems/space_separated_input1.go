package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func display_array(arr []int64, N int64) {

	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

}

func main() {
	var T int

	fmt.Scanf("%d\n", &T)

	for T != 0 {

		var N int64
		var arr []int64

		n, err := fmt.Scanf("%d\n", &N)

		if err != nil {
			fmt.Println(n, err)
		}

		reader := bufio.NewReader(os.Stdin)

		input_string, _ := reader.ReadString('\n')

		input_string = strings.TrimSpace(input_string) // TrimSpace function to remove leading and trailing whitespace as defined by Unicode.

		input := strings.Split(input_string, " ")

		fmt.Printf("%T\n", input)

		for _, value := range input {

			i, err := strconv.ParseInt(value, 8, 32)

			if err != nil {
				panic(err)
			}
			arr = append(arr, i)
		}

		display_array(arr, N)

		T--
	}

}
