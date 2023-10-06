package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	// logging
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	//profiling
	/*
		profiledAdd := getProfileOperation(add)
		profiledAdd(100, 200)

		profiledSubtract := getProfileOperation(subtract)
		profiledSubtract(100, 200)
	*/

	logAdd := getLogOperation(add)
	profiledLogAdd := getProfileOperation(logAdd)
	profiledLogAdd(100, 200)

	getProfileOperation(getLogOperation(subtract))(100, 200)
}

func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		op(x, y)
		log.Println("Operation completed")
	}
}

func getProfileOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started")
	op(x, y)
	log.Println("Operation completed")
}

/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

// 3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract  Result :", x-y)
}
