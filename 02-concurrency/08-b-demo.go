/* communication between goroutines */

package main

import (
	"fmt"
)

// consumer
func main() {
	ch := make(chan int)
	go add(ch, 100, 200)
	result := <-ch //receive
	fmt.Println(result)
}

// producer
func add(ch chan int, x, y int) {
	result := x + y
	ch <- result //send
}
