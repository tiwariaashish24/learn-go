package main

import "fmt"

func main() {

	fmt.Println("welcome to loops classes")

	days := []string{"Sunday", "monday", "Tuesday", "wednesday", "thusday", "friday", "saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// i dont want to print index so we use _ in that place lets check this

	// for _, day := range days {
	// 	fmt.Printf("the range %v \n", day)
	// }

	rougeValue := 10

	for rougeValue > 1 {
		fmt.Println(rougeValue)
		rougeValue--
	}

	for rougeValue < 20 {
		fmt.Println(rougeValue)
		rougeValue++
	}
}
