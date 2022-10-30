package main

import "fmt"

func reverseArray(a []int32) []int32 {

	arr := []int32{}

	for i := len(a) - 1; i >= 0; i-- {
		arr = append(arr, a[i])
	}

	return arr
}

func main() {

	a := []int32{1, 4, 3, 2}

	reversed := reverseArray(a)
	fmt.Println(reversed)
}
