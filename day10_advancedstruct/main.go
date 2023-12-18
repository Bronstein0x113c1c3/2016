package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}
type Supplier struct {
	name, city string
}

func (s *Supplier) PrintDetails() {
	fmt.Printf("Name: %v - City:%v  \n", s.name, s.city)
}
func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
func (p Product) PrintDetails() {
	fmt.Printf("Name: %v - Category: %v - Price: $%v \n", p.name, p.category, p.calcTax(0.2, 100))
}

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.price > threshold {
		return p.price * (rate + 1)
	}
	return p.price
}

func main() {
	// //creating a slice of products.
	// products := []*Product{
	// 	{"Kayak", "Watersports", 275},
	// 	{"Lifejacket", "Watersports", 48.95},
	// 	{"Soccer Ball", "Soccer", 19.50},
	// }
	// //collapse the whole command into smaller functions....
	// for _, p := range products {
	// 	// PrintDetails(p)
	// 	//or...
	// 	// p.PrintDetails()
	// 	//or
	// 	Product.PrintDetails(*p)
	// }

	//for supplier
	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.PrintDetails()
	}
}

func PrintDetails(p *Product) {
	fmt.Printf("Name: %v - Category: %v - Price: $%v \n", p.name, p.category, p.price)
}

//but, maybe, it's not enough...., the program is more compilated.....
