package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/gopxl/beep/flac"
	_ "github.com/gopxl/beep/flac"
	_ "github.com/gopxl/beep/mp3"
)

// func main() {
// 	f, err := os.Open("rightforyou.mp3")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	streamer, _, err := mp3.Decode(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	chan_sound := make(chan [][2]float64, 10)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)
// 	go func() {
// 		var streamer beep.Streamer = beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
// 			x := <-chan_sound
// 			for i := range samples {
// 				samples[i][0] = x[i][0]
// 				samples[i][1] = x[i][1]
// 			}
// 			return len(samples), true
// 		})
// 		sr := beep.SampleRate(44100)
// 		speaker.Init(sr, sr.N(time.Second))
// 		speaker.Play(streamer)
// 		wg.Done()
// 	}()
// 	for {
// 		result := make([][2]float64, 512)
// 		if _, ok := streamer.Stream(result); !ok {
// 			close(chan_sound)
// 			break
// 		}
// 		chan_sound <- result
// 	}
// 	wg.Wait()
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("someone get it......: %v \n", r.RemoteAddr)
		f, err := os.Open("04 - In My Blood.flac")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		streamer, _, err := flac.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()
		encoder := gob.NewEncoder(w)
		for {
			buffer := make([][2]float64, 512)
			if _, ok := streamer.Stream(buffer); !ok {
				return
			}
			encoder.Encode(&buffer)
		}
	})
	http.ListenAndServe(":8080", nil)
}
