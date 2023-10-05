/* concurrent safe state manipulation */
package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

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
	mutex.Lock()
	{
		counter += 1
		fmt.Printf("fn[%d], counter = %d\n", id, counter)
	}
	mutex.Unlock()

}
