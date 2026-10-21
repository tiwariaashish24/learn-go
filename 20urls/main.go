package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://aistudio.google.com/onboarding"

func main() {
	fmt.Println("welcome to url handling in goLang")
	fmt.Println(myUrl)

	//parsing url [extract informatin]

	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T", qparams)

	fmt.Println(qparams["coursename"]) //key value on url

	for _, val := range qparams {
		fmt.Println("The params is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "/tgolang",
		RawPath: "user=ashish",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}

/**
0. URL handling means parsing, creating, modifying, and extracting information from a URL
1. net/url is used to work with URLs.
2. url.Parse() breaks a URL into parts.
3. Scheme is the protocol, like https.
4. Host is the website name, like example.com.
5.  you can iterate the for loop and key value pair on it [url]
6.	you can construct the chunks and make the url
*/
