/* channel behaviors */
package main

import "fmt"

func main() {
	ch := make(chan int)

	/*
		data := <-ch
		ch <- 100
	*/

	/*
		ch <- 100
		data := <-ch

	*/

	/*
		go func() {
			ch <- 100
		}()

		data := <-ch
		fmt.Println(data)
	*/

	go func() {
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100
}
