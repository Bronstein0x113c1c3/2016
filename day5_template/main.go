package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	t, err := template.ParseFiles("template.txt")
	
	if err != nil {
		log.Println("Some problem in creating template.....")
		os.Exit(1)
	}
	lists := []struct {
		Name string
		ID   string
	}{
		{
			Name: "Bronstein",
			ID:   "1624640",
		},
		{
			Name: "Vladimir",
			ID:   "20218105",
		},
	}
	if err = t.ExecuteTemplate(os.Stdout, "template.txt", struct {
		List []struct {
			Name string
			ID   string
		}
	}{lists}); err != nil {
		log.Println("Cannot parse.....")
		os.Exit(1)
	}

}
