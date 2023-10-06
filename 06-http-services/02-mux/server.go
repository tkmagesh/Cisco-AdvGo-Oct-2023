package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
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

func getOneProductHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}

	if pid, ok := strconv.Atoi(id); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	} else {
		for _, product := range products {
			if product.Id == pid {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(product); err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintln(w, "Internal Server Error")
				}
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Product not found")
	}
}

// customers
func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/products/{id}", getOneProductHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", getProductsHandler).Methods(http.MethodGet)
	router.HandleFunc("/products", createProductHandler).Methods(http.MethodPost)
	router.HandleFunc("/customers", customersHandler)
	http.ListenAndServe(":8080", router)
}
