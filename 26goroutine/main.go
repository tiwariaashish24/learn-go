package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("hello")
	greeter("world")

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)

	}

}

/*
1. A goroutine is a lightweight task in Go.
2. Use the go keyword to start a goroutine.
3. Goroutines allow tasks to run concurrently.
4. The execution order of goroutines is not fixed.
5. time.Sleep() pauses execution for a given time.
6. The program exits when the main() function finishes.
7. Use sync.WaitGroup to wait for goroutines to finish.
*/
