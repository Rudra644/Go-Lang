package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices class")

	var fruitList = []string{"Tomato", "Apple", "Peach"}
	// fmt.Printf("The data type is: %T\n", fruitList)
	fmt.Println("The fruit list is: ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("The fruit list is: ", fruitList)

	fruitList = (fruitList[1:3])
	fmt.Println("Thr furit list is: ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 235
	highScores[1] = 345
	highScores[2] = 510
	highScores[3] = 978

	fmt.Println(highScores)

	highScores = append(highScores, 499, 107)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("The slice length is: ", len(highScores))
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"Java", "C++", "Python", "Ruby", "Rust"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
