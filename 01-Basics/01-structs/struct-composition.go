package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", p.Id, p.Name, p.Cost)
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

// Receive a pointer when the object state need to be changed
func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost - (p.Cost * (discountPercentage / 100))
}

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

// Overriding Product.Print() method
func (p PerishableProduct) Print() {
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f, Expiry = %q\n", p.Id, p.Name, p.Cost, p.Expiry)
}

func (p PerishableProduct) String() string {
	return fmt.Sprintf("%v, Expiry = %q", p.Product.String(), p.Expiry)
}

func main() {
	/*
		grapes := PerishableProduct{
			Product{100, "Grapes", 100},
			"2 Days",
		}
	*/

	grapes := PerishableProduct{
		Product: Product{Id: 100, Name: "Grapes", Cost: 100},
		Expiry:  "2 Days",
	}
	fmt.Printf("%#v\n", grapes)

	// Accessing the attributes of the composed type
	fmt.Println(grapes.Product.Id)
	fmt.Println(grapes.Id)

	// method inheritance
	/*
		grapes.Print()
		grapes.ApplyDiscount(10)
		grapes.Print()
	*/

	fmt.Println(grapes)
	grapes.ApplyDiscount(10)
	fmt.Println(grapes)

}
