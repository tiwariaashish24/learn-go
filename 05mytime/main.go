package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:03:04 Monday "))

	//create date by you

	createdDate := time.Date(2020, time.December, 12, 22, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006  Monday  15-04-05"))

}
