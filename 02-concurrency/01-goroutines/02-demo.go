package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1. When the main function completes, the scheduler will not wait for other goroutines to complete
2. the scheduler looks for other goroutines scheduled and attempts to start/resume their execution when the current function execution is blocked (for any reason)
*/

var wg sync.WaitGroup

func main() {
	defer func() {
		fmt.Println("[Deferred fn]")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Done")
	}()
	/*
		for i := 0; i < 10; i++ {
			wg.Add(1) // increment the wg counter by 1
			go f1()   // scheduling the execution of f1 through the scheduler
		}
	*/

	wg.Add(100) // increment the wg counter by 10
	for i := 0; i < 10; i++ {
		go f1() // scheduling the execution of f1 through the scheduler
	}
	f2()
	wg.Wait() // block the execution of this function until the wg counter becomes 0 (default value)
}

/*
func f1() {
	fmt.Println("f1 invoked")
}
*/

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
