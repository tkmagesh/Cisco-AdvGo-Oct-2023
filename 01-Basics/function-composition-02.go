package main

import (
	"fmt"
	"log"
	"time"
)

type Operation func(int, int)

type Middleware func(Operation) Operation

func composeMiddleware(middlewares ...Middleware) Middleware {
	return func(op Operation) Operation {
		for i := len(middlewares) - 1; i >= 0; i-- {
			op = middlewares[i](op)
		}
		return op
	}
}

func main() {

	/*
		logAdd := getLogOperation(add)
		profiledLogAdd := getProfileOperation(logAdd)
		profiledLogAdd(100, 200)

		getProfileOperation(getLogOperation(subtract))(100, 200)
	*/

	composedAdd := composeMiddleware(getLogOperation, getProfileOperation)(add)
	composedSubtract := composeMiddleware(getLogOperation, getProfileOperation)(subtract)

	composedAdd(100, 200)
	composedSubtract(100, 200)
}

func getLogOperation(op Operation) Operation {
	return func(x, y int) {
		log.Println("Operation started")
		op(x, y)
		log.Println("Operation completed")
	}
}

func getProfileOperation(op Operation) Operation {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func logOperation(op Operation, x, y int) {
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
