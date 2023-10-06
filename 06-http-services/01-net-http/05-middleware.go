package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// products
type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products = []Product{
	{101, "Pen", 10},
	{102, "Pencil", 5},
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Internal Server Error")
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Bad Request")
			return
		}
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Product created successfully")
	}
}

// customers
func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

// middlewares

func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func profileMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("Time taken:", elapsed)
	}
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func applyMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	for _, middleware := range m {
		h = middleware(h)
	}
	return h
}

func main() {

	// http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/", profileMiddleware(logMiddleware(indexHandler)))
	http.HandleFunc("/", applyMiddleware(indexHandler, logMiddleware, profileMiddleware))
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/customers", customersHandler)
	http.ListenAndServe(":8080", nil)
}
