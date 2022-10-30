package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	Barun := User{"Barun", "bhattacharjeebarun@gmail.com", true, 25}
	fmt.Printf("Barun details are: %+v\n", Barun)                       // +v for struc
	fmt.Printf("Name is %v and email is %v\n", Barun.Name, Barun.Email) // +v for struc

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
