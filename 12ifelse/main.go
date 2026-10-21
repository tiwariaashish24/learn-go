package main

import "fmt"

func main() {
	fmt.Println("welcome to if else /(control statement)")

	loginCount := 57
	var result string

	if loginCount < 10 {
		result = "Regular user"

	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "10 login count"
	}

	fmt.Println(result)

	if nums := 8; nums < 10 {
		fmt.Println("number is greater")
	} else {
		fmt.Println("nums is greater than no.")
	}

}
