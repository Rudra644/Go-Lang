package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "Hello World, This is a text file"

	file, err := os.Create("./SampleTextFile")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("The length of text file is: ", length)
	defer file.Close()
	ReadFile("./SampleTextFile")
}

func ReadFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is data of file: ", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
