package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to decode"))
	}
	fmt.Println(d.Name)
	buf := bytes.NewBuffer([]byte(d.Payload))
	base64decoder := base64.NewDecoder(base64.StdEncoding, buf)
	file, err := os.OpenFile("workdir/"+d.Name, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create file"))

		return
	}

	if _, err := io.Copy(file, base64decoder); err != nil {
		log.Printf("error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to write file"))

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successful writing file"))
	return

}
