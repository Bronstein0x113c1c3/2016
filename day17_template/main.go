package main

import (
	"os"
	"text/template"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)

}
func main() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemp := allTemplates.Lookup("template.html")
		err = Exec(selectedTemp)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

// for _, p := range Products {
// 	Printfln("Product: %v, Category: %v, Price: $%.2f",
// 		p.Name, p.Category, p.Price)
// }
