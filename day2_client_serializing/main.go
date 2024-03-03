package main

import (
	"encoding/gob"
	"log"
	"net"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Println(err)
		return
	}
	wg := &sync.WaitGroup{}
	encoder := gob.NewDecoder(conn)
	wg.Add(1)
	go func() {
		var name struct {
			Name string `json:"name"`
		}
		err := encoder.Decode(&name)
		log.Println(name)
		if err != nil {
			log.Println("Could not decode JSON.....")
		}
		conn.(*net.TCPConn).CloseRead()
		wg.Done()
	}()
	wg.Wait()
	// return
}
