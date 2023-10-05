/* Write a goroutine that generates the first 20 values in the fibonocci series */

package main

import "fmt"

func main() {
	nos := genFib()
	for no := range nos {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	fibCh := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < 20; i++ {
			fibCh <- x
			x, y = y, x+y
		}
		close(fibCh)
	}()
	return fibCh
}
