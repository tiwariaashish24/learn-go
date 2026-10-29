package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{} // add another RWMutex
	//mut := &sync.RWMutex{}

	var score = []int{0}

	fmt.Println("Race Condition")
	//funci not name body and run immeditaly

	// wg.Add(1) //you can direct aslo add the total add group in same time
	wg.Add(3) //here we add the final value in one time
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add()   ////no need to add when yu retrun the final add group in same time
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

	}(wg, mut)

	wg.Done()

	wg.Wait()
	fmt.Println(score)
}

/*
1. sync.WaitGroup waits for goroutines to finish.
2. wg.Add(3) sets the counter to 3.
3. wg.Done() decreases the counter by 1.
4. wg.Wait() waits until the counter becomes 0.
5. sync.Mutex protects shared data from concurrent access.
6. mut.Lock() locks the critical section.
7. mut.Unlock() releases the lock.
8. A goroutine is started using the go keyword.
9. An anonymous function is a function without a name.
10. A slice can be modified using append().
11. Use sync.RWMutex when you need separate read and write locks.
12. Always match each wg.Add(1) with a wg.Done().
13. Call wg.Add() before starting goroutines.
14. Do not call wg.Done() from main() to compensate for a missing Done().
*/
