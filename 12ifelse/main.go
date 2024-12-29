package main

import "fmt"

func main() {
	fmt.Println("Let's learn about if and else statements in go lang")

	TheNum := 10

	if TheNum > 10 {
		fmt.Println("The number is bigger than 10")
	} else if TheNum < 10 {
		fmt.Println("The number is less than 10")
	} else {
		fmt.Println("The number is equal to 10")
	}
}
