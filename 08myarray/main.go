package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Guava"
	fruitList[3] = "Plum"

	fmt.Println("The needed fruitlist is ", fruitList)
	fmt.Println("The needed fruitlist is ", len(fruitList))

	var vegiList = [4]string{"Potato", "tomato", "onion"}
	fmt.Println("vegi list is:", len(vegiList))
}
