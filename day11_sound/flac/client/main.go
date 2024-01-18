package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func buffer_stream() chan [][2]float64 {
	buffer_chan := make(chan [][2]float64)
	go func() {
		resp, err := http.Get("http://localhost:8080/")
		defer resp.Body.Close()
		if err != nil {
			close(buffer_chan)
			return
		}
		decoder := gob.NewDecoder(resp.Body)
		for {
			buffer := make([][2]float64, 512)
			if err := decoder.Decode(&buffer); err != nil {
				close(buffer_chan)
				return
			}
			buffer_chan <- buffer
		}
	}()
	return buffer_chan
	// resp,err:=
}
func newStreamer(buffer_chan chan [][2]float64, signal chan struct{}) beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		buffer := <-buffer_chan
		if len(buffer) == 0 {
			signal <- struct{}{}
			return 0, false

		} else {
			for i := range samples {
				samples[i][0] = buffer[i][0]
				samples[i][1] = buffer[i][1]
			}
			return len(samples), true
		}
	})
}
func main() {
	done := make(chan struct{})
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second*2))
	buffer_chan := buffer_stream()
	streamer := newStreamer(buffer_chan, done)
	speaker.Play(streamer)
	<-done
}
