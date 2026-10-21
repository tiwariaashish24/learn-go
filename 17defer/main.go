package main

import "fmt"

func main() {
	defer fmt.Println("Third")

	fmt.Println("First")
	fmt.Println("Second")

	defer fmt.Println("11")
	defer fmt.Println("21")
	defer fmt.Println("31")
	myDeffer()

}

func myDeffer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/**
1. defer work on last
2. in reverse  order is defered
3. they check and put into the last  so LIFo => last in first out check


Conclsion

:-- when you see defer you can put into the last.
*/
