/* communication between goroutines */

package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		var ch chan int
		ch = make(chan int)
	*/
	// combine declaration & initialization
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, ch, 100, 200)
	result := <-ch //receive
	wg.Wait()
	fmt.Println(result)
}

func add(wg *sync.WaitGroup, ch chan int, x, y int) {
	result := x + y
	ch <- result //send
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
