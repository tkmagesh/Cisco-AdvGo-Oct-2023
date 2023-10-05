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
	// explicit cancellation
	/*
		ctx, cancel := context.WithCancel(rootCtx)
		fmt.Println("Hit ENTER to stop...")
		go func() {
			fmt.Scanln()
			cancel()
		}()
	*/

	// time-based cancellation
	/*
		ctx, cancel := context.WithTimeout(rootCtx, time.Second*10)
		defer cancel()
	*/

	// combining explicit and time-based cancellation
	/*
		ctx, cancel := context.WithTimeout(rootCtx, time.Second*10)
		defer cancel()
		fmt.Println("Hit ENTER to stop...")
		go func() {
			fmt.Scanln()
			cancel()
		}()
	*/

	// Using contexts to send data
	cancelCtx, cancel := context.WithTimeout(rootCtx, time.Second*10)
	defer cancel()
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	valCtx := context.WithValue(cancelCtx, "key-1", "val-1")

	wg.Add(1)
	go fn(valCtx, wg)
	wg.Wait()
	fmt.Println("Done...!")
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("key-1 :", ctx.Value("key-1"))
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
