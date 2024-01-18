package main

import (
	"context"
	"encoding/gob"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

func main() {
	/*
		serializing project:
			-type of encoding: gob
			-server: send the gob
			-client: receive the gob


			-server will send the gob to client. client then decode and return to the server.....


			-sending, encoding in handle_conn.
	*/
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Println("Error creating listener.....")
		return
	}
	wg := &sync.WaitGroup{}
	signal_chan, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("The listener stopped receiving request....")
				return
			}
			wg.Add(1)
			go handle_conn(conn, wg)
		}
	}()
	<-signal_chan.Done()
	listener.Close()
	log.Println("Waiting for all connections to closed.....")
	wg.Wait()
}
func handle_conn(c net.Conn, wg *sync.WaitGroup) {
	name := &struct {
		Name string `json:"name"`
	}{"Bronstein"}
	encoder := gob.NewEncoder(c)
	encoder.Encode(name)
	c.Close()
	wg.Done()
}
