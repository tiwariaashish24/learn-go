package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to Slices")

	var fruitList = []string{"apple", "mango", "kiwi"}
	fmt.Println("the fruitList is:", fruitList)

	fruitList = append(fruitList, "plum", "litchi", "cherry")
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:3])

	highScore := make([]int, 4)

	highScore[0] = 789
	highScore[1] = 459
	highScore[2] = 365
	highScore[3] = 565
	// highScore[4] = 896   they shw error due to fix size of array so add append

	fmt.Println(highScore)

	highScore = append(highScore, 568, 855, 965, 364, 253)

	fmt.Println("the final highScore is :", highScore)
	fmt.Println("the final highScore is :", len(highScore))

	sort.Ints(highScore)
	fmt.Println("the sorted highScore is :", highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove the value frm slices

	var courses = []string{"javascript", "go", "typescript", "react", "nextjs"}
	fmt.Println(courses)

	courses = append(courses, "express", "sql")
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
