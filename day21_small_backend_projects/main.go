package main

import (
	"day21/api"
	"day21/config"
	"day21/db/mysql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// os.Setenv("httpmuxgo121", "0")
	host, port, username, password, dbname, db_port := config.Init()
	fmt.Println(host, port, username, password, dbname, db_port)
	db, err := mysql.Init_db(host, db_port, username, password, dbname)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	//init the database first
	/*
		if the database initiated successfully, go to next
		else -> stop the psogram...
	*/

	//init the handle from the database
	handlers := api.InitHandlers(&db)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /doughnuts", handlers.GetDoughnuts)
	mux.HandleFunc("GET /doughnuts/{d_type}/", http.HandlerFunc(handlers.GetDoughnutsWithType))
	mux.HandleFunc("POST /doughnuts", handlers.AddDoughnuts)
	//serve that :)
	//set up the server....
	log.Println("Serving time....")
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux); err != nil {
		panic(err)
	}
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)

		}
	}()
}
