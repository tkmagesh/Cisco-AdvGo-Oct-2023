/* Write a goroutine that generates the fibonocci series for 5 seconds
- refactoring the timeout logic
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	nos := genFib()
	for no := range nos {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(no)
	}
	fmt.Println("All the data received")
}

func genFib() <-chan int {
	timeOutCh := Timeout(7 * time.Second)
	fibCh := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-timeOutCh:
				fmt.Println("Timeout occurred")
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

func Timeout(d time.Duration) <-chan time.Time {
	timeOutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeOutCh <- time.Now()
	}()
	return timeOutCh
}
