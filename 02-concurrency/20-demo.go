/* Modify the program in such a way that it produces the fibonocci series until the user hits ENTER key
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	nos := genFib(stopCh)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for no := range nos {
		fmt.Println(no)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All the data received")
}

func genFib(stopCh chan struct{}) <-chan int {
	fibCh := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			default:
				fibCh <- x
				x, y = y, x+y
			}
		}
		close(fibCh)
	}()
	return fibCh
}
