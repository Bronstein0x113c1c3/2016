package main

import "encoding/base64"
func main(){
	base64.StdEncoding.	
}

/*
package main

import (
    "encoding/base64"
    "os"
    "io"
    "log"
)

func main() {
    // Open the file
    file, err := os.Open("yourfile.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Create a new base64 stream encoder
    encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)

    // Stream file to base64 encoder
    if _, err := io.Copy(encoder, file); err != nil {
        log.Fatal(err)
    }

    // Close the encoder
    if err := encoder.Close(); err != nil {
        log.Fatal(err)
    }
}
*/