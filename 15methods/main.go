package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	Barun := User{"Barun", "bhattacharjeebarun@gmail.com", true, 25}
	fmt.Printf("Barun details are: %+v\n", Barun) // +v for struct
	fmt.Printf("Name is %v and email is %v\n", Barun.Name, Barun.Email)

	Barun.GetStatus()
	Barun.NewMail()
	Barun.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
	fmt.Println("Email of this user is:", u.Email)
}

func (u User) NewMail() {
	u.Email = "test@go.dev" // it does not change the original object
	fmt.Println("Email of this user is:", u.Email)
}
