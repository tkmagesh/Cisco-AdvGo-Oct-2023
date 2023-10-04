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
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	//
	ch3 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		d3 := <-ch3
		fmt.Println("Ch3 data :", d3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case d1 := <-ch1:
			fmt.Println("Ch1 data :", d1)
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
		case d2 := <-ch2:
			fmt.Println("Ch2 data :", d2)
			/* default:
			fmt.Println("No channel operations are successful")
			*/
		}
	}

}
