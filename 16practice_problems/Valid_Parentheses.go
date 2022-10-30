package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValid(s string) bool {
	var stack = []string{}
	var top int = 0

	brackets := make(map[string]string)
	brackets["]"] = "["
	brackets["}"] = "{"
	brackets[")"] = "("

	for _, bracket := range s {

		first_bracket, ok := brackets[string(bracket)]
		if ok {
			if top == 0 {
				return false
			}
			if first_bracket != stack[top-1] {
				return false
			} else if first_bracket == stack[top-1] {
				stack = append(stack[:top-1])
				top--
			}

		} else {
			stack = append(stack, string(bracket))
			top++
		}

	}
	//fmt.Println(stack)
	return true && top == 0

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	var parentheses string = strings.TrimSpace(input)

	fmt.Println(isValid(parentheses))

}
