package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var count int
	wg := &sync.WaitGroup{}
	/*
		fmt.Println("runtime.NumCPU :", runtime.NumCPU())
		runtime.GOMAXPROCS(4)
	*/
	flag.IntVar(&count, "count", 0, "# of goroutines to spin")
	flag.Parse()
	fmt.Printf("Spinning %d goroutines... Hit ENTER to start\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Println("All Done... Hit ENTER to shutdown")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)

}
