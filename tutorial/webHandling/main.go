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

// model for courses - file
type Course struct {
	CourseId  string  `json:"courseid"`
	CouseName string  `json:"name"`
	Author    *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses []Course

//middlware, helper

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CouseName == ""
	return c.CouseName == ""

}

//Controllers

func main() {

}

// sever home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Bleh!</h1>"))
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

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No Data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data")
		return
	}

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse() {

}
