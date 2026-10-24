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
	fmt.Println("API-learncode")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs",
		CoursePrice: 499, Author: &Author{FullName: "munna bhaiya", Website: "http:munnabhaiya.dev"}}) //when pointer come you can pass the refrence
	courses = append(courses, Course{CourseId: "4", CourseName: "next-js",
		CoursePrice: 469, Author: &Author{FullName: "gudu-bhaiya", Website: "http:gudubhaiya.dev"}})

	//Router
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getSingelCourse).Methods("POST")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")

	//listen on port
	log.Fatal(http.ListenAndServe("4000", r))
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
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("content-type", "application/json")

	//grab id from req
	params := mux.Vars(r)
	//loop, id , remove , add with my id

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
	//TODO: send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	// fetch id
	params := mux.Vars(r)

	// loop, {id},  remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}

	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "No course found with given ID",
	})

}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all Courses")
	w.Header().Set("Content-Type", "application/json")

	courses = []Course{}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "All courses delete Successfully",
	})

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
16. w http.ResponseWriter — Used to send a response to the client, such as JSON, text, or HTML.
17. r *http.Request — Contains information about the incoming request, such as the URL, HTTP method, headers, and request body.
*/
