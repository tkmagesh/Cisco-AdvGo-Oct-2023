/* coordinating between multiple channels */

package main

import (
	"fmt"
	"sync"
	"time"
)

/* modify the program in such a way the data is consumed (printed) in the order in which they are produced */

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println(<-ch1)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println(<-ch2)
		wg.Done()
	}()

	wg.Wait()

}
