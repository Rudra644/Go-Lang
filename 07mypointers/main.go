package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers class Akka")

	// var ptr *int
	// fmt.Println("The value of pointer is: ", ptr)
	myNumber := 23

	var ptr = &myNumber

	fmt.Println("The value of pointer is: ", ptr)
	fmt.Println("The value of pointer is: ", *ptr)

	*ptr = *ptr + 2

	lol := *ptr + 1

	fmt.Println("The value of pointer is: ", *ptr)
	fmt.Println("The value of pointer is: ", lol)

}
