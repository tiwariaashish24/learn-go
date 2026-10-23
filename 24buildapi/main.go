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

//model for file

type Course struct {
	CourseId    string  `json: "courseid"`
	CourseName  string  `json: "coursename`
	CoursePrice int     `json: "price"`
	Author      *Author `json: "author"`
}

type Author struct {
	FullName string `json: "fullname`
	Website  string `json: "website"`
}

// create fake db
var courses []Course

//middleware, helper-func

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getSingelCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	//grab the data
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a course")
	w.Header().Set("content-type", "application/json")

	//what if: body  is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")

	}

	//what about-{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data is inside JSON")
		return
	}

	//generate unique id , string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

/*
1. Course struct stores course details.
2. Author struct stores author name and website.
3. JSON tags define field names in JSON data.
4. *Author is a pointer to an Author struct.
5. courses is a slice used as a fake database.
6. mux.NewRouter() creates a router for API routes.
7. ServeHome() handles the home page request.
8. getAllCourses() returns all courses as JSON.
9. mux.Vars(r) gets URL parameters.
10. params["id"] gets the course ID from the URL.
11. json.NewEncoder(w).Encode() sends data as JSON.
12. createOneCourse() handles requests to create a course.
13. r.Body contains data sent in the request.
14. Check errors when decoding JSON or reading request data.
15. Return HTTP 404 when a course is not found.
*/
