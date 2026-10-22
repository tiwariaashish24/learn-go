package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb golang")
	// PerformGetRequest()
	// performPostJsonRequest()
	PerformPostFromRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/GET"

	response, err := http.Get(myUrl)
	HandleError(err)
	fmt.Print(response)
	fmt.Println("Status code:", response.StatusCode)

	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	// Store the response in a strings.Builder
	var responseString strings.Builder
	contentt, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(contentt)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func performPostJsonRequest() {
	const myUri = "http://localhost:3000/items"

	requestBody := strings.NewReader(`
	
	{
		"courseName": "Learn Golang",
		"price": 2500,
		"country": "india",
		"Language": "English"

	}
	
	`)

	response, err := http.Post(myUri, "application/json", requestBody)
	HandleError(err)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(content)
	fmt.Println(string(content))

	// fmt.Println("the all data is", response)
	// defer response.Body.Close()
}

// func PerformPostFromRequest() {
// 	const myUrl = "http://localhost:3000/itemsform"

// 	data := url.Values{}

// 	data.Add("firstname", "ashish")
// 	data.Add("middlename", "kumar")
// 	data.Add("lastname", "tiwari")

// 	response, err := http.PostForm(myUrl, data)
// 	HandleError(err)

// 	content, err := ioutil.ReadAll(response.Body)
// 	HandleError(err)
// 	fmt.Println(string(content))

// 	defer response.Body.Close()

// }

func PerformPostFromRequest() {
	const myUrl = "http://localhost:3000/itemsform"

	data := url.Values{}
	data.Add("firstname", "ashish")
	data.Add("middlename", "kumar")
	data.Add("lastname", "tiwari")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Response:", string(content))
}

/*
1. net/http is used to send web requests in Go.
2. GET is used to receive data from a server.
3. POST is used to send data to a server.
4. http.Get() sends a GET request.
5. http.Post() sends a POST request.
6. strings.NewReader() converts a string into a reader.
7. application/json tells the server that we are sending JSON data.
8. ioutil.ReadAll() reads the response body into bytes.
9. string(content) converts bytes into readable text.
10. defer response.Body.Close() closes the response body when the function ends.
11. response.StatusCode gives the HTTP status code.
12. HandleError() checks errors and panics if an error occurs.
13. localhost means your own computer.
14. Port 3000 is where your local server receives requests in this example.
15. A response body should be read only once.
16. Use io.ReadAll() instead of ioutil.ReadAll() in modern Go.
*/
