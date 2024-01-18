package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		// storeTotal += group.TotalPrice(category)
		go group.TotalPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		storeTotal += <-channel
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}
func (group ProductGroup) TotalPrice(category string, result_chan chan float64) (total float64) {
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	defer func() { result_chan <- total }()
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
