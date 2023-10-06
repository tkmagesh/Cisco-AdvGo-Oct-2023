package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
type AppServer struct {
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (appServer *AppServer) AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.routes[route] = handler
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := r.URL.Path
	if handler, ok := appServer.routes[route]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not Found")
		return
	} else {
		handler(w, r)
	}
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}
*/

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

func main() {
	/*
		appServer := &AppServer{}
		appServer.AddRoute("/", indexHandler)
		appServer.AddRoute("/products", productsHandler)
		appServer.AddRoute("/customers", customersHandler)
	*/
	appServer := http.NewServeMux()
	appServer.HandleFunc("/", indexHandler)
	appServer.HandleFunc("/products", productsHandler)
	appServer.HandleFunc("/customers", customersHandler)
	http.ListenAndServe(":8080", appServer)
}
