package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch and case learn")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough //generate next case automaticaly
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and role dice again ")

	default:
		fmt.Println("you won the game")

	}
}
