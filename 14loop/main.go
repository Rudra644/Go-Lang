package main

import "fmt"

func main() {
	fmt.Println("Let's learn about loop in golang")

	Days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}
	fmt.Println(Days)

	// for i := 0; i < len(Days); i++ {
	// 	fmt.Println(Days[i])
	// }

	// for i := range Days {
	// 	fmt.Println(Days[i])
	// }

	// for index, day := range Days {
	// 	fmt.Printf("The index is %v and the value is %v\n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 4 {
			goto nf
		}

		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println(rougeValue)
		rougeValue++
	}

nf:
	fmt.Println("Jump to NodesForest.com")
}
