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

// Models
type Course struct {
	CourseName  string  `json:"courseName"`
	CourseId    string  `json:"courseId"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middlewares/ helpers
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

//Controllers

// Home Controller
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Home Page of Nikhil's Website</h1>"))
}

// Get All Courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses from DB")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)

}

// Get One Course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course based on Id")

	w.Header().Set("Content-Type", "application/json")

	//grab the course id
	params := mux.Vars(r)

	//loop through each course
	//Find matching Id and return that course
	//If not found, return a msg

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with given ID found")
	return
}

// Create one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")

	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is empty, I want Data!!!")
		return
	}

	var course Course
	//take json and put in object
	_ = json.NewDecoder(r.Body).Decode(&course)

	//what if: json passed in body is empty, {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Json Is empty...")
		return
	}

	//generate unique Id for a new course
	//Convert that to string
	//append that to list of courses (database)

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func main() {

}
