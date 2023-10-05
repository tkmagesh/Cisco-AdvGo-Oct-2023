/* Errors from goroutines */
package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	doneCh, errCh := fn(true) // pass true or false to control the outcome of the function "fn"

	select {
	case status := <-doneCh:
		fmt.Println("Status :", status)
	case err := <-errCh:
		fmt.Println("error : ", err)
	}
}

func fn(err bool) (<-chan bool, <-chan error) {
	doneCh := make(chan bool)
	errCh := make(chan error, 1) // by using a buffered channel we are making receiving the error an operation activity
	go func() {
		time.Sleep(2 * time.Second)
		if err {
			errCh <- errors.New("unknown error")
			doneCh <- false
		} else {
			doneCh <- true
		}
	}()
	return doneCh, errCh
}
