package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []courses{
		{"ReactJs Bootstrap", 299, "Youtube", "123abc", []string{"web", "frontend"}},
		{"MERN Bootstrap", 199, "Youtube", "157abc", []string{"full-stack", "web-dev"}},
		{"Express Bootstrap", 299, "Youtube", "abc123", nil},
	}

	// package this json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	JsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJs Bootstrap",
                "Price": 299,
                "website": "Youtube",
                "tags": ["web","frontend"]
    }
	`)

	var lcoCourses courses

	CheckValid := json.Valid(JsonDataFromWeb)

	if CheckValid {
		fmt.Println("Json is valid")
		json.Unmarshal(JsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Incase you want to have key value pair from JSON

	var myOnlineJson map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &myOnlineJson)
	fmt.Printf("%#v\n", myOnlineJson)

	for k, v := range myOnlineJson {
		fmt.Printf("For key %v the value is %v and the type is %T \n", k, v, v)
	}
}
