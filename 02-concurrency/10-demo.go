/* streaming data through channels */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
	}()
	return ch
}
