package main

import (
	"encoding/json"
	"net/http"
)

func HandleJsonRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//enfore the json content-type
	json.NewEncoder(writer).Encode(Products)
}
func init() {
	http.HandleFunc("/json", HandleJsonRequest)
}
