package main

import (
	"io"
	"log"
	"text/template"
)

var all_templ *template.Template

func init() {
	var err error
	funcMap := template.FuncMap{
		"checkGPA": checkGPA,
	}
	all_templ, err = template.New("student_templates").Funcs(funcMap).ParseGlob("template/*_with_board.txt")
	if err != nil {
		log.Fatalln("Error loading template")
	}
	log.Println("Tempalte initiated")
}
func PrintStudent(student Student, w io.Writer) {
	all_templ.Lookup("student_with_board").Execute(w, student)
}
func PrintStudentList(list []Student, w io.Writer) {
	all_templ.Lookup("main_with_board").Execute(w, list)

}
