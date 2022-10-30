package main

import "fmt"

func display_array(arr []int32, N int32) {

	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

}

func main() {
	var T int32

	fmt.Scanf("%d\n", &T)

	for T != 0 {

		var N int32
		n, err := fmt.Scanf("%d\n", &N)

		if err != nil {
			fmt.Println(n, err)
		}

		arr := make([]int32, N)

		var i int32 = 0

		for ; i < N; i++ {
			if _, err := fmt.Scan(&arr[i]); err != nil {
				panic(err)
			}
		}

		display_array(arr, N)

		T--
	}
}
