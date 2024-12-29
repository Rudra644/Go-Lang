package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests in golang")
	PerformPostFormRequests()
}

func PerformGetRequests() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	// var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	// byteCount, _ := responseString.Write(content)

	// fmt.Println(byteCount)
	// fmt.Println(responseString.String())

	data := string(content)
	fmt.Println("The response is: ", data)
}

func PerformPostRequests() {

	const myurl = "http://localhost:8000/post"

	// Fake Json Payload

	requestBody := strings.NewReader(`
	{
	"coursename":"Let's go with go lang",
	"Price":0,
	"platform":"youtube"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response from server is: ", string(content))
}

func PerformPostFormRequests() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("name", "Ankit")
	data.Add("last", "Chandwada")
	data.Add("email", "Ankitchandwada007@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
