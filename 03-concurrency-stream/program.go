package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}

	dataCh := make(chan int)
	oddCh := make(chan int)
	evenCh := make(chan int)

	fileWg.Add(1)
	go Source("data1.dat", dataCh, fileWg)
	fileWg.Add(1)
	go Source("data2.dat", dataCh, fileWg)

	processWg.Add(1)
	go Splitter(dataCh, evenCh, oddCh, processWg)

	processWg.Add(1)
	go Sum("Even-Total.txt", evenCh, processWg)

	processWg.Add(1)
	go Sum("Odd-Total.txt", oddCh, processWg)

	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
}

func Source(fileName string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if no, err := strconv.Atoi(val); err == nil {
			ch <- no
		}
	}
}

func Splitter(ch chan int, evenCh chan int, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(evenCh)
	defer close(oddCh)
	counter := 0
	for data := range ch {
		counter += 1
		if data%2 == 0 {
			evenCh <- data
		} else {
			oddCh <- data
		}
	}
	fmt.Println("counter :", counter)
}

func Sum(title string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	total := 0
	count := 0
	for val := range ch {
		total += val
		count += 1
	}
	fmt.Println(title, "total :", total, "count :", count)

	// refactor the following into the Merger() function so that both even sum & odd sum are written to the same file

	file, err := os.Create(title)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "Total : %d\n", total)
}
