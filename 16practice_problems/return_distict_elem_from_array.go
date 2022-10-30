package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func set(arr []int) []int {

	var counter = [10000]int{}
	var set = []int{}

	for _, value := range arr {

		counter[value]++

		if counter[value] == 1 {
			set = append(set, value)
		}
	}

	return set

}

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	for T != 0 {

		var N int
		var arr = []int{}
		fmt.Scanf("%d\n", &N)

		reader := bufio.NewReader(os.Stdin)

		input_string, _ := reader.ReadString('\n') // "1 2 3 4 5\r\n"

		input_string = strings.TrimSpace(input_string) // "1 2 3 4 5"

		input := strings.Split(input_string, " ") // ['1' '2' '3' '4' '5']

		for _, value := range input {
			i, err := strconv.ParseInt(value, 16, 32)
			if err != nil {
				panic(err)
			}

			arr = append(arr, int(i))
		}

		distinct_arr := set(arr)
		fmt.Println(distinct_arr)

		T--
	}
}
