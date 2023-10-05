package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go fn(cancelCtx, wg)
	wg.Wait()
	fmt.Println("Done...!")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancellation signal received...")
			break LOOP
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now())
		}
	}
}
