package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	for _, currentCourse := range courses {
		if currentCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Duplicate course names not allowed")
			return
		}
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

// Update a course
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")

	w.Header().Set("Content-Type", "application/json")

	//grab id fromm request

	params := mux.Vars(r)

	//add validations as required

	//loop through courses, find that Id, remove that course, add new course with received Id

	for index, course := range courses {
		if params["id"] == course.CourseId {
			//removing at specific index
			courses = append(courses[:index], courses[index+1:]...)
			//get updated values
			var updatedCourse Course
			json.NewDecoder(r.Body).Decode(&updatedCourse)
			//set same Id for updated values
			updatedCourse.CourseId = params["id"]
			//add updated course in list
			courses = append(courses, updatedCourse)

			json.NewEncoder(w).Encode(updatedCourse)
			return

		}
	}

	//if Id not found in list

	json.NewEncoder(w).Encode("No course found for this course id")
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, find Id, remove that course

	for index, course := range courses {
		if course.CourseId == params["id"] {
			//remove at specfic index
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course removed successfully")
			break
		}
	}

	json.NewEncoder(w).Encode("Course with that id not found")
}

func main() {

	fmt.Println("Nikhil's Backend - Building APIs")
	fmt.Println("http://localhost:8000/")

	r := mux.NewRouter()

	//seeding - filling data

	courses = append(courses, Course{
		CourseId:    "111",
		CourseName:  "Psychology",
		CoursePrice: 599,
		Author:      &Author{FullName: "Abhishek More", Website: "https://google.com"},
	})

	courses = append(courses, Course{
		CourseId:    "222",
		CourseName:  "Economics",
		CoursePrice: 799,
		Author:      &Author{FullName: "Rehan Shaikh", Website: "https://netflix.com"},
	})

	courses = append(courses, Course{
		CourseId:    "333",
		CourseName:  "Comedy",
		CoursePrice: 699,
		Author:      &Author{FullName: "Rupesh Bodkhe", Website: "https://yahoo.com"},
	})

	courses = append(courses, Course{
		CourseId:    "444",
		CourseName:  "Computer",
		CoursePrice: 999,
		Author:      &Author{FullName: "Shubham Rane", Website: "https://facebook.com"},
	})

	//handle routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listen and serve
	log.Fatal(http.ListenAndServe(":8000", r))

}
