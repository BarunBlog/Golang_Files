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

// Model for course -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"` // - to not show it on json
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB table
var tbl_course []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func (c *Course) IsUnique() bool {
	for _, value := range tbl_course {
		if value.CourseName == c.CourseName {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("API creating")

	r := mux.NewRouter()

	// seeding
	tbl_course = append(tbl_course, Course{
		CourseId:    "2",
		CourseName:  "ReactJS",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Barun",
			Website:  "lco.dev",
		},
	})

	tbl_course = append(tbl_course, Course{
		CourseId:    "4",
		CourseName:  "MERN Stack",
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Barun",
			Website:  "lco.dev",
		},
	})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/create", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API tutorial</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // setting headers
	json.NewEncoder(w).Encode(tbl_course)              // encode the data into json format
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) // extrating all the params coming from the requests

	// loop through courses, find matching id return response
	for _, course := range tbl_course {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decoded json value will be stored into course variable
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	if !course.IsUnique() {
		json.NewEncoder(w).Encode("Please provide unique course name")
		return
	}

	// genarate unique id, string
	// append course into courses

	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano()) // unique number
	course.CourseId = strconv.Itoa(rand.Intn(100))

	tbl_course = append(tbl_course, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range tbl_course {
		if course.CourseId == params["id"] {
			tbl_course = append(tbl_course[:index], tbl_course[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			tbl_course = append(tbl_course, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range tbl_course {
		if course.CourseId == params["id"] {
			tbl_course = append(tbl_course[:index], tbl_course[index+1:]...)

			json.NewEncoder(w).Encode("Successfully deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("Id not found")
}
