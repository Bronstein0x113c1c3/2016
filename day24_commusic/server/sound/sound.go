package sound

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gopxl/beep/flac"
	"github.com/gopxl/beep/mp3"
)

type Audio struct {
	dir string
}

func NewAudioInstance(dir string) *Audio {
	return &Audio{dir: dir}
}

func (a *Audio) ListAllSong() ([]Song, error) {
	files, err := ioutil.ReadDir(a.dir)
	if err != nil {
		return []Song{}, err
	}
	var songlist []Song
	id := 1
	for _, file := range files {
		if strings.Contains(file.Name(), ".flac") {
			songlist = append(songlist, GetInfoSongFLAC(a.dir, file.Name(), id))
			id += 1
		} else if strings.Contains(file.Name(), ".mp3") {
			songlist = append(songlist, GetInfoSongMP3(a.dir, file.Name(), id))
			id += 1
		}
	}
	return songlist, nil

}
func (a *Audio) AudioDataStream(song_name string) (chan [][2]float64, error) {
	if strings.Contains(song_name, ".mp3") {
		return a.AudioDataStreamMP3(song_name)
	} else if strings.Contains(song_name, ".flac") {
		return a.AudioDataStreamFLAC(song_name)
	} else {
		return nil, errors.New("Cannot decode!!!")
	}
}
func (a *Audio) AudioDataStreamMP3(song_name string) (chan [][2]float64, error) {

	//open the song first, make it into a streamer
	f, err := os.Open(a.dir + song_name)
	log.Printf("Opening: %v \n", a.dir+song_name)
	if err != nil {
		return nil, err
	}

	streamer, _, err := mp3.Decode(f)
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
func (a *Audio) AudioDataStreamFLAC(song_name string) (chan [][2]float64, error) {

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
