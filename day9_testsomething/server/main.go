package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name    string `json:"Name"`
	Payload string `json:"Payload"`
}

func main() {
	http.DefaultServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bonjour!"))
	})
	http.DefaultServeMux.HandleFunc("/send", ReceiveFile)
	log.Println("initiated")
	http.ListenAndServe(":8080", nil)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	log.Println("something is coming.....")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	d := &Data{} //create a data object
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		fmt.Println(err)
	}
	fmt.Println(*d)

}
