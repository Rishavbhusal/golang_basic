package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//model for course -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware/helper -file
func (c Course) IsEmpty() bool {
	return c.CourseName == ""

}

func main() {
	fmt.Println("Api ")
	r := mux.NewRouter()
	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 200, Author: &Author{Fullname: "Rishav Bhusal", Website: "go.len"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Mern stack", CoursePrice: 323, Author: &Author{Fullname: " Rishav Bhusal", Website: "Mern.ln"}})

	// routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/onecourse/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/create", createOnePost).Methods("POST")
	r.HandleFunc("/update/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}", deleteOneCourse).Methods("DELETE")
	// listening to port
	log.Fatal(http.ListenAndServe(":2000", r))
}

// controllers  -file

//server home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Api </h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "appilcation/json")
	// grab id from request
	params := mux.Vars(r)
	// fmt.Printf("The datatype is %T", params)

	// loop through courses, find matching id and retturn response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found from given id")
	return

}

func createOnePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating course")
	w.Header().Set("Content-Type", "application/json")

	// when body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// when {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("There is no data inside")
		return
	}

	for _, title := range courses {
		if title.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Name Cant be same")
			return
		}

	}

	// generate a unique id ,string
	// append course into courses
	course.CourseId = strconv.Itoa(rand.IntN(100) + 1)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updade course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req

	params := mux.Vars(r)

	// loop,id,,remove,add with paramsid
	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)

			return
		}

	}
	// send a response when id is not found

	// if course.CourseId != params["id"] {
	// 	json.NewEncoder(w).Encode("ID not found")
	// }

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	// get the id
	params := mux.Vars(r)

	// loop,findid,remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// json.NewEncoder(w).Encode("Deleted successfully")
			break
		}
		// 	json.NewEncoder(w).Encode("Bad request")
		// 	return
	}
}
