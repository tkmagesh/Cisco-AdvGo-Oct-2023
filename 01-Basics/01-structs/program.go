package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	/*
		var product struct {
			Id   int
			Name string
			Cost float32
		}
		product.Id = 100
		product.Name = "Pen"
		product.Cost = 10
	*/
	// fmt.Println(product)
	/*
		var product struct {
			Id   int
			Name string
			Cost float32
		} = struct {
			Id   int
			Name string
			Cost float32
		}{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	var product Product = Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Printf("%#v\n", product)
	// PrintProduct(product)
	product.Print()

	/*
		var obj struct{}
		fmt.Println(obj)
	*/

	// structs are values
	p2 := product //create a copy of product
	p2.Cost = 20
	fmt.Println("product :", product)
	fmt.Println("p2 :", p2)

	// use pointers to create references
	productPtr := &product
	// fmt.Println(productPtr.Id, productPtr.Name, productPtr.Cost)
	productPtr.Print()

	productPtr.ApplyDiscount(10)
	productPtr.Print()
}

/*
func PrintProduct(p struct {
	Id   int
	Name string
	Cost float32
}) {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", p.Id, p.Name, p.Cost)
}
*/

/*
func PrintProduct(p Product) {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", p.Id, p.Name, p.Cost)
}
*/

// The above function as a method
func (p Product) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", p.Id, p.Name, p.Cost)
}

// Receive a pointer when the object state need to be changed
func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost - (p.Cost * (discountPercentage / 100))
}
