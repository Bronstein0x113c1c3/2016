package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	in "server/input"

	"github.com/gordonklaus/portaudio"
)

func main() {
	input, err := in.New()
	if err != nil {
		log.Fatal(err)
	}
	// input.Start()

	// io.Pipe()
	// data_chan := make(chan []float32)

	var s string
	fmt.Print("Waiting...: ")
	fmt.Scanln(&s)
	fmt.Println("Waiting...: ")
	input.Stop()
	// // input.Terminate()
	// binary.Read
	// for {
	// 	buf := make([]int16, 8196)
	// 	binary.Read(input.GetStream(), binary.LittleEndian)
	// }
	portaudio.Initialize()
	defer portaudio.Terminate()
	in := bytes.NewReader(input.GetStream().(*bytes.Buffer).Bytes())
	buf := make([]int16, 8196)
	reading, err := portaudio.OpenDefaultStream(0, 1, 16000, len(buf), &buf)
	reading.Start()
	if err != nil {
		log.Fatalln(err)
	}
	signal := make(chan struct{})
	go func() {
		for {
			err := binary.Read(in, binary.LittleEndian, &buf)
			reading.Write()
			if err != nil {
				signal <- struct{}{}
				return
			}
		}
	}()
	<-signal
}

//caller and callee!!!
