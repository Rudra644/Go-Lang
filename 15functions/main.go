package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	resultN := adder(4, 9)
	fmt.Println("The sum is ", resultN)

}

func adder(Val1 int, Val2 int) int {
	return Val1 + Val2
}
