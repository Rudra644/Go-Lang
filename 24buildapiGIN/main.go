package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (c *Course) IsEmpty() bool {
	return c.CourseName == "" // You can add more checks if needed
}

// fake db
var courses []Course

func main() {
	// Initialize Gin router
	r := gin.Default()

	seedData()

	// Routes
	r.GET("/", serveHome)
	r.GET("/courses", getAllCourses)
	r.GET("/course/:id", getOneCourse)
	r.POST("/course", createOnceCourse)
	r.PUT("/course/:id", updateOneCourse)

	// Listen to a port
	r.Run(":4000")

}

// seeding fake data
func seedData() {
	courses = append(courses, Course{CourseId: "2", CourseName: "React JS", CoursePrice: 299, Author: &Author{Fullname: "Ankit Chandwada", Website: "React.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Go lang", CoursePrice: 249, Author: &Author{Fullname: "Rohit Sharma", Website: "go.dev"}})
}

func serveHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Test API built using Gin!")
}

func getAllCourses(c *gin.Context) {
	c.JSON(http.StatusOK, courses)
}

func getOneCourse(c *gin.Context) {
	id := c.Param("id")

	for _, course := range courses {
		if course.CourseId == id {
			c.JSON(http.StatusOK, course)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Course not found for " + id + " ID"})
}

func createOnceCourse(c *gin.Context) {

	var course Course

	// Prase JSON body
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong JSON format"})
		return
	}

	// Validate the data
	if course.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course body cannot be empty"})
		return
	}

	// Check for duplicate data

	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Course name already exists"})
			return
		}
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	c.JSON(http.StatusCreated, course)
}

func updateOneCourse(c *gin.Context) {
	id := c.Param("id")

	var updatedCourse Course

	// Prase JSON body
	if err := c.ShouldBindJSON(&updatedCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong JSON format"})
		return
	}

	// Validate the data
	if updatedCourse.IsEmpty() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course body cannot be empty"})
		return
	}

	// Check for duplicate course name

	for _, existingCourse := range courses {
		if existingCourse.CourseName == updatedCourse.CourseName && existingCourse.CourseId != id {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Course name already exists"})
			return
		}
	}

	for i, course := range courses {
		if course.CourseId == id {
			updatedCourse.CourseId = id // Don't update the original ID
			courses[i] = updatedCourse
			c.JSON(http.StatusOK, updatedCourse)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Course ID " + id + " not found"})
}
