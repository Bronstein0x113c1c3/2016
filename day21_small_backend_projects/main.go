package main

import (
	"day21/db/mysql"
	"fmt"
)

func main() {
	db, err := mysql.Init_db("localhost", 3306, "root", "0x113c1c3", "greg_list")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic("Shiet, some problem....")
	} else {
		println("Done")
	}
	fmt.Println(db.GetDoughnuts())
	//init the database first
	/*
		if the database initiated successfully, go to next
		else -> stop the program...
	*/

	//init the handle from the database

	//serve that :)

}
