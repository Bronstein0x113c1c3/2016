package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("Bonjour")
	template_string := "Hello, {{.Name}} \n"
	t.Parse(template_string)
	t.Execute(os.Stdout, struct {
		Name string
	}{"Jonathan"})
}
