package main

import "fmt"

func main() {

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("This is list of fruits: ", fruitList)
	fmt.Println("This is list of fruits: ", len(fruitList))

	var veggyList = [3]string{"Tomato", "Chilli", "Ginger"}

	fmt.Println("This is list of fruits: ", veggyList)
	fmt.Println("This is list of fruits: ", len(veggyList))
}
