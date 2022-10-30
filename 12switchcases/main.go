package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // will generate number between 1 and 5 + 1

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough // as soon as the value is 3, this case will run and also the next case
	case 4:
		fmt.Println("Dice value is 4")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")

	default:
		fmt.Println("Whops, what was that!")

	}
}
