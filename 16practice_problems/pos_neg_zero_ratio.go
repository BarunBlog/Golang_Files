package main

import "fmt"

func plusMinus(arr []int32) {
	var pos_counter float32 = 0
	var neg_counter float32 = 0
	var zero_counter float32 = 0

	for _, value := range arr {
		if value > 0 {
			pos_counter++
		} else if value < 0 {
			neg_counter++
		} else {
			zero_counter++
		}
	}
	size := float32(len(arr))

	fmt.Printf("%v\n%v\n%v\n", pos_counter/size, neg_counter/size, zero_counter/size)

}

func main() {
	arr := []int32{1, 3, 4, -1, -5, 0, 4}
	plusMinus(arr)
}
