package main

import (
	"day20/ifaces"
	"fmt"
)

func main() {
	r := ifaces.Result{
		Animal: ifaces.Dog{
			Name:  "X",
			Breed: "Y",
		},
	}
	fmt.Println(r.Speak())
}
