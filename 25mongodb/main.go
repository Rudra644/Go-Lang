package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rudra644/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to MongoDB")
	r := router.Router()
	fmt.Print("Server is getting started... \n")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Listening at port 8080")

}
