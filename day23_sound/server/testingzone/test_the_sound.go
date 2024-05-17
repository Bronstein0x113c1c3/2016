package testingzone

import (
	"fmt"
	"log"
	"os"
	"server/sound"
	"time"

	"github.com/dhowden/tag"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

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
func Run_the_test() {
	a := sound.NewAudioInstance("./testingzone/")
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second*5))
	audio_stream, err := a.AudioDataStream("11 - Why.flac")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed")
	}
	done := make(chan struct{})
	streamer := newStreamer(audio_stream, done)
	speaker.Play(streamer)
	<-done
}

func Readinginfo(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	meta, err := tag.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}
	res := ""
	res += fmt.Sprintf("name: %v \n", meta.Title())
	res += fmt.Sprintf("artist name: %v \n", meta.Artist())
	res += fmt.Sprintf("album name: %v \n", meta.Album())
	i, _ := meta.Track()
	res += fmt.Sprintf("track: %v \n", i)
	res += fmt.Sprintf("published year: %v", meta.Year())
	fmt.Println(res)
}
