// package main

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	"github.com/gordonklaus/portaudio"
// )

// const sampleRate = 44100
// const seconds = 2

// func main() {
// 	portaudio.Initialize()
// 	defer portaudio.Terminate()
// 	buffer := make([]int16, 8192)

// 	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, len(buffer), func(out []int16) {
// 		resp, err := http.Get("http://localhost:8080/audio")
// 		chk(err)
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		responseReader := bytes.NewReader(body)
// 		binary.Read(responseReader, binary.LittleEndian, &buffer)
// 		for i := range out {
// 			out[i] = buffer[i]
// 		}
// 	})
// 	chk(err)
// 	chk(stream.Start())
// 	time.Sleep(time.Second * 40)
// 	chk(stream.Stop())
// 	defer stream.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }

//	func chk(err error) {
//		if err != nil {
//			panic(err)
//		}
//	}

package main

import (
	"encoding/binary"
	"io"
	"net/http"

	"github.com/gordonklaus/portaudio"
)

func main() {
	resp, err := http.Get("http://localhost:8080/mp3")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	portaudio.Initialize()
	defer portaudio.Terminate()

	buffer := make([]int16, 1024)                                                 // Increase the buffer size
	stream, err := portaudio.OpenDefaultStream(0, 2, 44100, len(buffer), &buffer) // Increase the sample rate
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	stream.Start()
	defer stream.Stop()
	done := make(chan struct{})
	go func() {
		for {
			err = binary.Read(resp.Body, binary.LittleEndian, buffer)
			if err == io.EOF {
				done <- struct{}{}
				break
			} else if err != nil {
				panic(err)
			}

			stream.Write()
		}
	}()
	<-done
}
