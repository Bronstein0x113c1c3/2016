package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}
type Supplier struct {
	name, city string
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
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

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func (p Product) getName() string {
	return p.name
}
func (p Product) getCost(recur bool) float64 {
	return p.price
}
func (s Service) getName() string {
	return s.description
}
func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

type Account struct {
	accNum   int
	expenses []Expense
}

func main() {
	// //creating a slice of products.
	// products := ProductList([]Product{
	// 	{"Kayak", "Watersports", 275},
	// 	{"Lifejacket", "Watersports", 48.95},
	// 	{"Soccer Ball", "Soccer", 19.50},
	// })
	// kayak := Product{"Kayak", "Watersports", 275}
	// insurance := Service{"Boat Cover", 12, 89.50}
	// fmt.Printf("Product: %v, Price: %v \n", kayak.name, kayak.price)
	// fmt.Printf("Service: %v, Price: %v \n", insurance.description, insurance.monthlyFee*float64(insurance.durationMonths))

	// for category, total := range products.calcCategoryTotals() {
	// 	fmt.Printf("Category: %v, Total: %v \n", category, total)
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
	// suppliers := []*Supplier{
	// 	{"Acme Co", "New York City"},
	// 	{"BoatCo", "Chicago"},
	// }
	// for _, s := range suppliers {
	// 	s.PrintDetails()
	// }
	// expenses := []Expense{
	// 	Product{"Kayak", "Watersports", 275},
	// 	Service{"Boat Cover", 12, 89.50},
	// }
	// for _, expense := range expenses {
	// 	fmt.Printf("Expense: %v, Cost: %v \n", expense.getName(), expense.getCost(true))
	// }
	// fmt.Println("Total:", calcTotal(expenses))
	account := Account{
		accNum: 12345,
		expenses: []Expense{
			newProduct("Kayak", "Watersports", 275),
			Service{"Boat Cover", 12, 89.50},
		},
	}
	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total: ", calcTotal(account.expenses))

}

func PrintDetails(p *Product) {
	fmt.Printf("Name: %v - Category: %v - Price: $%v \n", p.name, p.category, p.price)
}

//but, maybe, it's not enough...., the program is more compilated.....
