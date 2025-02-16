package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	websiteList := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://github.com",
		"https://golang.org",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(website string) {
	defer wg.Done()
	result, err := http.Get(website)

	if err != nil {
		fmt.Println("Their is an error in website")
	}

	fmt.Printf("The status code for %v is %v \n", website, result.StatusCode)
}
