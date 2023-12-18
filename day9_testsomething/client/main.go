package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("workdir/Invitation-to-Computer-Science-G.-Michael-Schneider.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffer to capture the output
	var buf bytes.Buffer

	// Create a new base64 stream encoder that writes to the buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)

	// Stream file to base64 encoder
	if _, err := io.Copy(encoder, file); err != nil {
		log.Fatal(err)
	}

	// Close the encoder
	if err := encoder.Close(); err != nil {
		log.Fatal(err)
	}

	// log.Println(buf.String())

	//so, the drama is coming...

	type Data struct {
		Name    string `json:"Name"`
		Payload string `json:"Payload"`
	}
	d := &Data{
		Name:    "Invitation-to-Computer-Science-G.-Michael-Schneider.pdf",
		Payload: buf.String(),
	}
	pr, pw := io.Pipe()

	// Create a new JSON encoder that writes to the pipe writer
	go func() {
		enc := json.NewEncoder(pw)
		if err := enc.Encode(d); err != nil {
			log.Fatal(err)
		}
		pw.Close()
	}()

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "http://localhost:8080/send", pr)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status
	log.Println("response Status:", resp.Body.)
}
