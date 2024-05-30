package main

import (
	"io"
	"log"
	"net"

	"github.com/gordonklaus/portaudio"
	// "golang.org/x/mobile/exp/sprite/portable"
)

func serve(c net.Conn, data_chan chan []byte) {

}
func chan_to_stream(data_chan chan []byte) {
}
func main() {
	data_chan := make(chan []byte)
	// buf := make([]byte, 1024)
	buffer := make([]int16, 1024)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	ios, outs := io.Pipe()
	portaudio.Initialize()
	input, err := portaudio.OpenDefaultStream(0, 2, 44100, len(buf))

	for {
		conn, err := lis.Accept()
		go serve(conn, data_chan)
	}
}
