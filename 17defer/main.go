package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	Numb0()
}

func Numb0() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
