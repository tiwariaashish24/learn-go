package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://aistudio.google.com/onboarding"

func main() {
	fmt.Println("Welcome to  LCO web request in golang")

	response, err := http.Get(url) //give the http request on url
	ErrorCheck(err)

	fmt.Printf("the response is %T \n", response)

	databytes, err := ioutil.ReadAll(response.Body)
	ErrorCheck(err)
	// fmt.Println(databytes)         //original data bytes
	// fmt.Println(string(databytes)) //databytes conver

	content := string(databytes)
	fmt.Println(content)

	response.Body.Close() //caller responsibility to close the connection
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

/**
1. you not revcie any copy of response actually geting a response
2. caller responsibility to close the connection always put in the last in fn
*/
