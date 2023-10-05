/* concurrent safe state manipulation */
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count += 1
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println(counter.count)
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
