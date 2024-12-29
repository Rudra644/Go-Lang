package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.stroustrup.com/index.html"

func main() {
	fmt.Println("Learn web requests in go lang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The response is type of %T: ", response)

	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	data := string(databyte)

	fmt.Println(data)
}
