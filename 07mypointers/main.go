package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers") //pointer is just a direct memory refrence

	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 69

	var ptr = &myNumber
	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("My new Value is: ", myNumber)
	fmt.Println("My new Value is: ", *ptr)

}
