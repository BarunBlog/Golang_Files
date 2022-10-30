package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list len is:", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}
	fmt.Printf("Vegy list type is: %T\n", vegList)
	fmt.Println("Vegy list is:", vegList)
	fmt.Println("Veg list len is:", len(vegList))

	fmt.Println(vegList[1:])

}
