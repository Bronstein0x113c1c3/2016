package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:1807")
	if err != nil {
		log.Fatal("cannot call to server")
	}
	
}
