package main

import (
	"fmt"
	"log"
	"mongodb/router"
	"net/http"
)

func main() {
	fmt.Println("Server is getting started...")

	r := router.Router()

	fmt.Println("Listening on port 4001")
	log.Fatal(http.ListenAndServe(":4001", r))
}
