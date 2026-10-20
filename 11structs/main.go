package main

import "fmt"

func main() {
	fmt.Println("welcome to struct")
	//no inheretace in golang  and No super not parents

	ashish := User{"Aashish", "tiwari@go.in", true, 19}
	fmt.Println(ashish)

	fmt.Printf("aashish details are %+v\n", ashish) //+ add they give more details
	fmt.Printf("Name is %v and Email is %v.", ashish.Name, ashish.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
