package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func handle_conn(c net.Conn, c chan os.Signal)
func main() {
	signal_chan := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	endpoint, _ := net.ResolveTCPAddr("tcp4", ":1807")
	listener, _ := net.ListenTCP("tcp4", endpoint)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
		}
		handle_conn(conn)
	}

}
