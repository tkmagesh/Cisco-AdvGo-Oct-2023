/* concurrent safe state manipulation */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)

}
