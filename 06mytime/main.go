package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the study of time")

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006"))
}
