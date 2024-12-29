package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Lets learn switch and case in golang")

	source := rand.NewSource(time.Now().UnixNano())
	diceNumber := rand.New(source).Intn(6) + 1
	fmt.Println("Your dice number is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1 please open your goti")
	case 2:
		fmt.Println("You got 2 please move your goti")
	case 3:
		fmt.Println("You got 3 please move your goti")
	case 4:
		fmt.Println("You got 4 please move your goti")
		fallthrough
	case 5:
		fmt.Println("You got 5 please move your goti")
	case 6:
		fmt.Println("Hurry!! You got 6 please roll again")
	default:
		fmt.Println("Out of scope")
	}
}
