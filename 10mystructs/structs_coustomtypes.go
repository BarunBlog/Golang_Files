package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func changeX(pt *Point) {
	pt.x = 10
}

func changeX_(pt Point) {
	pt.x = 10
}

type Circle struct {
	radius float64
	//center *Point // embed the Point struct.
	*Point //If you don't want to give it a name, instead embed the point to the struct
}

func main() {
	var p1 Point = Point{1, 2}
	var p2 Point = Point{-5, 7}

	fmt.Println(p1.x, p1.y)
	fmt.Println(p2.x, p2.y)

	p1.x = 7

	//var p3 Point = Point{x: 3} // set x to three, y left to default value 0
	p3 := &Point{y: 3} // p3 is equal to pointer to the Point{y:3}
	fmt.Println(p3)
	changeX(p3) // Changing the property x of the object p3
	fmt.Println(p3)

	p4 := Point{y: 3} // p3 is equal to pointer to the Point{y:3}
	fmt.Println(p4)
	changeX_(p4) // Changing the property x of the object p3
	fmt.Println(p4)

	// Embeded struct
	//point := &Point{y:3}
	//c1 := Circle{4.56, point}

	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1)
	//fmt.Println(c1.center)
	//fmt.Println(c1.center.x)
	//c1.center.x = 5
	//fmt.Println(c1.center.x)
	fmt.Println(c1.x)
}
