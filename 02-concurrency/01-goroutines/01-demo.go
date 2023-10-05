package main

import (
	"fmt"
	"time"
)

/*
1. When the main function completes, the scheduler will not wait for other goroutines to complete
2. the scheduler looks for other goroutines scheduled and attempts to start/resume their execution when the current function execution is blocked (for any reason)
*/
func main() {
	go f1() // scheduling the execution of f1 through the scheduler
	f2()

	// time.Sleep(4 * time.Second) // blocking the executin of the main function so that the scheduler can look for other goroutines scheduled and execute them

	fmt.Scanln() //blocking
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
}

func f2() {
	fmt.Println("f2 invoked")
}
