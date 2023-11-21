package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:0x113c1c3@tcp(localhost:3306)/testing")
	defer db.Close()
	if err != nil {
		log.Panicln("Got problem: ", err)
	} else {
		log.Println("Connected")
		// db.Close()
	}
	type Info struct {
		Name string
		ID   int
	}
	// prepare, err := db.Prepare("INSERT INTO student_info VALUES (?,?)")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // defer prepare.Close()
	// for i := 0; i < 100; i++ {
	// 	_, _ = prepare.Exec(i, strconv.Itoa(i))
	// }
	// prepare.Close()
	rows, err := db.Query("SELECT * FROM student_info")
	defer rows.Close()
	for rows.Next() {
		var student Info
		rows.Scan(&student.ID, &student.Name)
		fmt.Println(student)
	}
}
