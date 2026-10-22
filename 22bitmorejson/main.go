package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name       string `json: "name"`
	Price      int
	CourseName string `json: "playlist"`
	Website    string `json: "platform"`
	UserName   string `json: "username"`
	Password   string `json:"-"`
	Type       []string
}

func main() {

	fmt.Println("Welcome to JSON -> original data convert into json")
	// EncodedJson()
	DecodeJsonData()

}

func EncodedJson() {

	myCourse := []course{

		{"Ashish", 2500, "Learn Golang", "https://go.dev", "ashish", "demo123", []string{"Go", "Backend"}},
		{"Rahul", 3000, "Learn Java", "https://java.com", "rahul", "demo456", []string{"Java", "DSA"}},
		{"Aman", 2000, "Learn Python", "https://python.org", "aman", "demo789", []string{"Python", "AI"}},
		{"Priya", 3500, "Learn JavaScript", "https://developer.mozilla.org", "priya", "demo101", []string{"JavaScript", "Web Development"}},
	}

	finalResult, err := json.MarshalIndent(myCourse, "", "\t")

	// finalResult, err := json.Marshal(myCourse)

	if err != nil {
		panic(err)
	}

	fmt.Printf(string(finalResult))
}

func DecodeJsonData() {
	jsonDataFromWeb := []byte(`
	  {
                "Name": "Priya",
                "Price": 3500,
                "CourseName": "Learn JavaScript",
                "Website": "https://developer.mozilla.org",
                "UserName": "priya",
                "Type": ["JavaScript", "Web Development"]
        }
	`)
	var locCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &locCourse)

		fmt.Printf("%#v\n", locCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where you just wnat to add data to key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	//2nd way

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and Value is %v and Type is %T \n", k, v, v)
	}

}

/*
1. json.Marshal() converts Go data into JSON bytes.
2. json.MarshalIndent()--> (data, prefix, indent)converts Go data into formatted JSON bytes.
3. string(finalResult) converts JSON bytes into a string.
4. fmt.Println() prints the result.
5. JSON only includes exported struct fields (fields starting with a capital letter).
6. json.Unmarshal() converts JSON data into Go data.
*/
