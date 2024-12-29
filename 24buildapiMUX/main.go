package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func (c Course) IsEmpty() bool {
	return c.CourseName == ""
}

// seeding fake data
var courses []Course

func main() {
	// Initialize Mux router
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "React JS", CoursePrice: 299, Author: &Author{Fullname: "Ankit Chandwada", Website: "React.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Go lang", CoursePrice: 249, Author: &Author{Fullname: "Rohit Sharma", Website: "go.dev"}})

	// Define the routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Start the server
	fmt.Println("Starting the server on port 4000")
	http.ListenAndServe(":4000", r)

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Test API built using MUX!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	fmt.Printf(" Params is of type %T and its value is %v \n", params, params)

	var course Course
	for _, course = range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found for " + params["id"] + " ID in database")
}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// handel no data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// handel empty JSON {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("JSON can not be empty")
		return
	}

	// create random ID for new data
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	// Generate a random integer between 1 and 100 (inclusive)
	courseID := generator.Intn(100) + 1
	course.CourseId = strconv.Itoa(courseID)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, delete, add and update id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder((r.Body)).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, delete (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found for " + params["id"] + " ID in database")
}
