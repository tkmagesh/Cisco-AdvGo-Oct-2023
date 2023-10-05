/* streaming data through channels */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	// count := rand.Intn(50)
	count := 10
	go func() {
		for i := 0; i < count; i++ {
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
