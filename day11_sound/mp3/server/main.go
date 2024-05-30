// package main

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/bobertlo/go-mpg123/mpg123"
// )

// func main() {
// 	/*
// 		plan:
// 			1. decode the mp3 file first.
// 			2. info of mp3 file? sample_rate, channel, buffer_length,.....
// 			3. send the buffer to the client.

// 		at client:
// 			1. decode the buffer with portaudio, with all of these info.
// 			2. from the buffer, encode it to ......
// 			3. hear that.
// 	*/
// 	if len(os.Args) < 2 {
// 		fmt.Println("missing required argument:  input file name")
// 		return
// 	}
// 	fileName := os.Args[1]
// 	//use file name.....
// 	decoder, err := mpg123.NewDecoder("")
// 	chk(err)
// 	chk(decoder.Open(fileName))
// 	//maybe server closed, but.....
// 	defer decoder.Close()
// 	//get info....
// 	rate, channel, encoding := decoder.GetFormat()
// 	//reconfig
// 	decoder.FormatNone()
// 	log.Printf("info of the mp3 file: rate: %v, channel: %v, encoding: %v", rate, channel, encoding)
// 	decoder.Format(rate, channel, mpg123.ENC_16)

// 	//done the preparation

// 	http.HandleFunc("/audio", func(w http.ResponseWriter, r *http.Request) {
// 		out := make([]int16, 8192)
// 		flusher, ok := w.(http.Flusher)
// 		if !ok {
// 			panic("expected http.ResponseWriter to be an http.Flusher")
// 		}

// 		w.Header().Set("Connection", "Keep-Alive")
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("X-Content-Type-Options", "nosniff")
// 		w.Header().Set("Transfer-Encoding", "chunked")
// 		w.Header().Set("Content-Type", "audio/wave")
// 		for {
// 			audio := make([]byte, 2*len(out))
// 			_, err = decoder.Read(audio)
// 			if err == mpg123.EOF {
// 				break
// 			}
// 			chk(err)
// 			chk(binary.Read(bytes.NewBuffer(audio), binary.BigEndian, out))
// 			binary.Write(w, binary.LittleEndian, &out)
// 			flusher.Flush()
// 		}
// 	})
// 	http.ListenAndServe(":8080", nil)

// }
// func chk(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// /*
// flusher, ok := w.(http.Flusher)
// 		if !ok {
// 			panic("expected http.ResponseWriter to be an http.Flusher")
// 		}

// w.Header().Set("Connection", "Keep-Alive")
// w.Header().Set("Access-Control-Allow-Origin", "*")
// w.Header().Set("X-Content-Type-Options", "nosniff")
// w.Header().Set("Transfer-Encoding", "chunked")
// w.Header().Set("Content-Type", "audio/wave")
//
//	for true {
//		binary.Write(w, binary.BigEndian, &buffer)
//		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
//		return
//	}*/

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kvark128/minimp3"
	// "github.com/sdobz/go-mpg123"
	// "github.com/sdobz/go-mpg123"
)

func main() {
	http.HandleFunc("/mp3", func(w http.ResponseWriter, r *http.Request) {
		// decoder, err := mpg123.NewDecoder("")
		// chk(err)
		// chk(decoder.Open("rightforyou.mp3"))
		file, err := os.Open("rightforyou.mp3")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		mp3dec := minimp3.NewDecoder(file)
		//
		// decoder.Format()
		// decoder.Format(rate, channel, mpg123.ENC_16)
		buf := make([]byte, 256)
		for {
			if _, err := mp3dec.Read(buf); err != nil {
				break
			}
			w.Write(buf)

		}
		log.Printf("%v: done \n", r.RemoteAddr)
	})
	http.ListenAndServe(":8080", nil)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

// package main

// import (
// 	"net/http"
// 	"os"

// 	"github.com/hajimehoshi/go-mp3"
// )

// func main() {
// 	http.HandleFunc("/mp3", func(w http.ResponseWriter, r *http.Request) {
// 		f, _ := os.Open("rightforyou.mp3")
// 		defer f.Close()

// 		d, _ := mp3.NewDecoder(f)
// 		// defer d.Close()

// 		data := make([]byte, 256) // Smaller buffer size
// 		for {
// 			if _, err := d.Read(data); err != nil {
// 				break
// 			}

// 			w.Write(data)
// 		}
// 	})

// 	http.ListenAndServe(":8080", nil)
// }

//little
