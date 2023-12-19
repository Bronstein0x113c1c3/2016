package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func main() {
	f, err := os.Open("rightforyou.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	chan_sound := make(chan [][2]float64)
	go func() {
		var streamer beep.Streamer = beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
			x := <-chan_sound
			for i := range samples {
				samples[i][0] = x[i][0]
				samples[i][1] = x[i][1]
			}
			return len(samples), true
		})
		sr := beep.SampleRate(48000)
		speaker.Init(sr, sr.N(time.Second*3))
		speaker.Play(streamer)
	}()
	for {
		result := make([][2]float64, 512)
		if _, ok := streamer.Stream(result); !ok {
			close(chan_sound)
			break
		}
		chan_sound <- result
	}

}
