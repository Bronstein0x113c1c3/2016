package sound

import (
	"log"
	"os"

	"github.com/gopxl/beep/flac"
)

type Audio struct {
	dir string
}

func NewAudioInstance(dir string) *Audio {
	return &Audio{dir: dir}
}

func (a *Audio) ListAllSong() {

}
func (a *Audio) AudioDataStream(song_name string) (chan [][2]float64, error) {

	//open the song first, make it into a streamer
	f, err := os.Open(a.dir + song_name)
	log.Printf("Opening: %v \n", a.dir+song_name)
	if err != nil {
		return nil, err
	}

	streamer, _, err := flac.Decode(f)
	if err != nil {
		return nil, err
	}
	data_chan := make(chan [][2]float64)
	go func() {
		for {
			samples := make([][2]float64, 512)
			if _, ok := streamer.Stream(samples); !ok {
				close(data_chan)
				return
			}
			data_chan <- samples
		}
	}()
	// through the streamer, push everything into a channel
	//return that channel
	return data_chan, nil
}
