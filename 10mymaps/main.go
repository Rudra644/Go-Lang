package main

import "fmt"

func main() {
	fmt.Println("These are maps in go lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("These are the languages ", languages)
	fmt.Println("Full form of JS is: ", languages["JS"])

	delete(languages, "PY")
	fmt.Println("These are the languages ", languages)

	for key, value := range languages {
		fmt.Printf("For key the %v and its valaue is %v \n", key, value)
	}

}
