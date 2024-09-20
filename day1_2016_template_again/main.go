package main

import (
	"log"
	"os"
	"text/template"
)

type StudentInfo struct {
	Name   string
	ID     int
	School string
	Gender string
	// BloodType byte
}

func main() {
	student_list := []StudentInfo{
		{Name: "Emily Johnson",
			ID:     123456,
			School: "Lincoln High School",
			Gender: "Female",
		},
		{Name: "Michael Smith",
			ID:     789012,
			School: "Roosevelt Middle School",
		},
		{Name: "Sarah Williams",
			ID:     345678,
			School: "Jefferson Elementary",
		},
		{Name: "David Brown",
			ID:     901234,
			School: "Washington High School",
		},
		{Name: "Jessica Garcia",
			ID:     567890,
			School: "Kennedy Academy",
		}}
	templ, err := template.ParseFiles("template.txt")
	if err != nil {
		log.Printf("cannot parse the template: %v \n", err)
		return
	}
	_ = templ.Execute(os.Stdout, student_list)
	// if err != nil {
	// 	log.Println(err)
	// }
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"text/template"
// )

// type Person struct {
// 	FullName string
// 	Age      int
// }

// func main() {
// 	people := []Person{
// 		{"Emily Johnson", 29},
// 		{"Michael Smith", 34},
// 		{"Sarah Williams", 22},
// 		{"David Brown", 41},
// 		{"Jessica Garcia", 27},
// 	}

// 	tmpl := `
// |----------------------|-----|
// {{- range .People }}
// | {{ .FullName | printf "%-20s" }} | {{ .Age | printf "%-3d" }} |
// {{- end -}}`

// 	t, err := template.New("table").Parse(tmpl)
// 	if err != nil {
// 		fmt.Println("Error creating template:", err)
// 		return
// 	}

// 	// Execute the template with the data
// 	err = t.Execute(os.Stdout, struct {
// 		FullName string
// 		Age      string
// 		People   []Person
// 	}{
// 		FullName: "Full Name",
// 		Age:      "Age",
// 		People:   people,
// 	})
// 	if err != nil {
// 		fmt.Println("Error executing template:", err)
// 		return
// 	}
// }
