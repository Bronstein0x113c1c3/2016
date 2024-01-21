package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	all_temp, err := template.ParseGlob("template/*.txt")
	if err != nil {
		log.Fatal("Cannot load template")
	}
	all_temp.Lookup("main_with_board").Execute(os.Stdout, list_of_student)

}
