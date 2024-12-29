package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://nodesforest.com/dashboard?coin=HootChain&currency=USD"

func main() {
	fmt.Println("Welcome to handeling URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("The result of type qparams is: %T \n", qparams)

	fmt.Println(qparams["coin"])

	for key, value := range qparams {
		fmt.Printf("%v: %v\n", key, value)
	}

	partsofurl := &url.URL{
		Scheme:   "https",
		Host:     "nodesforest",
		Path:     "/user",
		RawQuery: "admin=yes",
	}

	fmt.Println(partsofurl)
}
