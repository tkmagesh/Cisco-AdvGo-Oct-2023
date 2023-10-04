/* communication between goroutines */

package main

import (
	"fmt"
	"sync"
)

// Communicate by sharing memory
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200)
	wg.Wait()
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, x, y int) {
	result = x + y
	wg.Done()
}

/*
func main() {
	wg := &sync.WaitGroup{}
	var result int
	wg.Add(1)
	go add(wg, 100, 200, &result)
	wg.Wait()
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, x, y int, result *int) {
	*result = x + y
	wg.Done()
}
*/
