package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTIme := time.Now()
	fmt.Println(presentTIme)

	fmt.Println(presentTIme.Format("01-02-2006 15:04:05 Monday")) // standered format // also 02-01-2006 works

	createdDate := time.Date(2021, time.October, 25, 15, 47, 12, 10, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
