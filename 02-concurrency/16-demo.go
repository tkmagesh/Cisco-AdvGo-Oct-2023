/* coordinating between multiple channels */

package main

import (
	"fmt"
	"time"
)

/* modify the program in such a way the data is consumed (printed) in the order in which they are produced */

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch1 <- i * 2
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(300 * time.Millisecond)
			ch2 <- (i * 2) + 1
		}
		close(ch2)
	}()

	//
	ch3 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		d3 := <-ch3
		fmt.Println("Ch3 data :", d3)
	}()

	var opCompleted int
	for opCompleted < 3 {
		select {
		case d1, isOpen := <-ch1:
			fmt.Println("Ch1 data :", d1)
			if !isOpen {
				opCompleted += 1
			}
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
			opCompleted += 1
		case d2, isOpen := <-ch2:
			fmt.Println("Ch2 data :", d2)
			if !isOpen {
				opCompleted += 1
			}
			/* default:
			fmt.Println("No channel operations are successful")
			*/
		}
	}
	fmt.Println("Done")

}
