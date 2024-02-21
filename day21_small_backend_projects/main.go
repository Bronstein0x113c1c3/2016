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
	host, port, username, password, dbname, db_port := config.Init()
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
	mux := http.DefaultServeMux
	mux.HandleFunc("/get/", http.HandlerFunc(handlers.GetDoughnuts))
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
