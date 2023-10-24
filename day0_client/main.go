package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1807")
	defer conn.Close()
	if err != nil {
		log.Fatal("cannot call to server")
	}
	go func(conn net.Conn) {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
	}(conn)
}
