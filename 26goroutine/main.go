package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //pointer

func main() {
	go greeter("hello")
	greeter("world")
	websitelist := []string{
		"https://loc.dev",
		"https://google.com",
		"https://go.dev",
		"https://github.com",
		"https://facebook.com",
		"https://netflix.in",
	}

	for _, web := range websitelist {

		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)

	}

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Println("%d status code for %s", res.StatusCode, endpoint)
}

/*
1. A goroutine is a lightweight task in Go.
2. Use the go keyword to start a goroutine.
3. Goroutines allow tasks to run concurrently.
4. The execution order of goroutines is not fixed.
5. time.Sleep() pauses execution for a given time.
6. The program exits when the main() function finishes.
7. Use sync.WaitGroup to wait for goroutines to finish.
8. how to wait group utilised ?
9. wg helps main() wait for goroutines to finish.
10. wg.Add(1) increases the WaitGroup counter.
12. go starts a function in a goroutine.
13. defer wg.Done() decreases the counter when the function finishes.
14. wg.Wait() waits until the counter becomes zero.
15. Always call Add() before starting the goroutine.
16. Check HTTP errors before accessing the response.
:--------Remember: Add() → go function() → Done() → Wait().----------------:

*/
