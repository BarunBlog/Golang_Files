package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 200
}

func main() {
	var x []int = []int{3, 4, 5} // slice

	y := x // y and x pointing to the same slice
	y[0] = 100
	//x[0] = 101

	fmt.Println(x, y)
	changeFirst(x)
	fmt.Println(x, y)

	var mymap_x map[string]int = map[string]int{"hello": 3} // map
	mymap_y := mymap_x

	mymap_y["y"] = 100
	mymap_y["x"] = 200
	fmt.Println(mymap_x, mymap_y)

	var myx [2]int = [2]int{3, 4}
	myy := myx // does not point to the same array
	myy[0] = 100

	fmt.Println(myx, myy)
}
