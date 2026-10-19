package main

import "fmt"

const publicToken = "hiiiformaashish"

func main() {
	var username string = "aashish"
	fmt.Println(username)
	fmt.Printf("Variable is a type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is a type: %T \n", isLoggedIn)

	var isSmall uint8 = 255
	fmt.Println(isSmall)
	fmt.Printf("Variable is a type: %T \n", isSmall)

	var isFloat float32 = 255.009089
	fmt.Println(isFloat)
	fmt.Printf("Variable is a type: %T \n", isFloat)

	//check default value and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is a type: %T \n", anotherVariable)

	var anotherUser string
	fmt.Println(anotherUser)
	fmt.Printf("Variable is a type: %T \n", anotherUser)

	//implicit type

	var movieName = "mirzapur the movie"
	fmt.Println(movieName) //they can trat as string
	fmt.Printf("Variable is a type: %T \n", movieName)

	totalCount := 658 // := declear a new variable && use inside functions, but not directly at package level
	fmt.Println(totalCount)
	fmt.Printf("Variable is a type: %T \n", totalCount)

	fmt.Println(publicToken)
	fmt.Printf("Variable is a type: %T \n", publicToken)
}
