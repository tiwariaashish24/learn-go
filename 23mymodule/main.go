package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":4000", r) /// for running the server
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<H1> Welcome to golang playlist  by Hitesh chaudhary</H1>"))
}

/*
1. gorilla/mux is used to handle routes in Go.
2. mux.NewRouter() creates a new router.
3. HandleFunc() connects a URL path to a function.
4. Methods("GET") allows GET requests only.
5. http.ListenAndServe(":4000", r) starts the server on port 4000.
6. log.Fatal() prints an error and stops the program if the server fails.
7. serveHome() handles requests to the home page.
8. http.ResponseWriter sends a response to the client.
9. *http.Request contains information about the incoming request.
10. w.Write() sends data or HTML to the browser.
11. go mod tidy cleans up project dependencies.
*/
