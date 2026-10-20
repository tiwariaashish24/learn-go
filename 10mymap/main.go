package main

import "fmt"

func main() {
	fmt.Println("Welcome to Map ")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["j"] = "java"
	languages["cpp"] = "cplus"
	languages["c"] = "c"
	languages["ru"] = "ruby"

	fmt.Println(languages)

	delete(languages, "c")
	fmt.Println(languages)

	//loops are interesting

	for _, value := range languages {
		fmt.Printf("for key v, value is %v\n ", value)
	}

}
