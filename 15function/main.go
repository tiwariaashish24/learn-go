package main

import "fmt"

func main() {
	fmt.Println("Welocme to functin in golang")

	gretter()

	result := 25
	fmt.Println(result)

	output := adder(3, 5)

	fmt.Println("the total sum = ", output)

	proRes, _ := proAdder(2, 4, 6, 8, 9) //the two type of method returnign is
	fmt.Println(proRes)

	proRes, myMessage := proAdder(2, 4, 6, 8, 9)
	fmt.Println(proRes, myMessage)
}

//function signature you can add the data type outside the value [int] so this is value signature
func adder(valOne int, valTwo int) int {
	return valOne + valTwo

}

//pro fn
func proAdder(values ...int) (int, string) { //... dots are verodioc fn and you retun two value with comman in pernthesis
	total := 0
	for _, val := range values {
		total += val

	}
	return total, "this is the final sum"
}

func gretter() {
	fmt.Println("hello whats your name ")
}
