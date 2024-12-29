package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is structs in golang")

	ankit := User{"Ankit", "Ankit@gmail.com", true, 23}

	fmt.Println("This is user data :", ankit)
	fmt.Printf("The data for user is: %+v\n", ankit)
	fmt.Printf("The name is %v and his email is %v.\n", ankit.Name, ankit.Email)
	ankit.GetStatus()
	ankit.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("The email for this user is: ", u.Email)
}
