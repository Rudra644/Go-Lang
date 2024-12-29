package main

import "fmt"

func main() {
	fmt.Println("This is structs in golang")

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	ankit := User{"Ankit", "Ankit@gmail.com", true, 23}
	fmt.Println("This is user data :", ankit)
	fmt.Printf("The data for user is: %+v\n", ankit)
	fmt.Printf("The name is %v and his email is %v.", ankit.Name, ankit.Email)

}
