/* communication between goroutines */

package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := add(100, 200)
	result := <-ch //receive
	fmt.Println(result)
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		result := x + y
		ch <- result //send
	}()
	return ch
}
