package main

import (
	"encoding/binary"
	"log"
	"os"
	"sync"

	"github.com/gordonklaus/portaudio"
	"github.com/kvark128/minimp3"
	"gopkg.in/hraban/opus.v2"
	// "github.com/pion/pkg/"
)

const (
	DefaultBitDepth       = 16
	DefaultChannels       = 2
	DefaultSampleRate     = 48000
	DefaultFrameSize      = 480
	DefaultOpusDataLength = 1000
)

func mp3_to_opus(filename string, data_chan chan []byte, wg *sync.WaitGroup) {
	encoder, err := opus.NewEncoder(DefaultSampleRate, DefaultChannels, opus.AppAudio)
	if err != nil {
		log.Println("Cannot set the encoder!!!")
		return
	}
	mp3_file, err := os.Open("305.mp3")
	if err != nil {
		log.Println("cannot open file!!!!!")
	}
	mp3_decoder := minimp3.NewDecoder(mp3_file)
	defer wg.Done()
	defer close(data_chan)
	for {
		input := make([]int16, DefaultFrameSize)
		output := make([]byte, 3000)
		err := binary.Read(mp3_decoder, binary.LittleEndian, input)
		if err != nil {
			return
		}
		n, _ := encoder.Encode(input, output)
		data_chan <- output[:n]
	}
}
func opus_to_mp3(data_chan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	dec, err := opus.NewDecoder(DefaultSampleRate, DefaultChannels)
	if err != nil {
		log.Println("cannot ....")
		return
	}
	output := make([]int16, DefaultFrameSize)
	portaudio.Initialize()
	defer portaudio.Terminate()
	streamer, err := portaudio.OpenDefaultStream(0, DefaultChannels, DefaultSampleRate, DefaultFrameSize, &output)
	streamer.Start()
	defer streamer.Stop()
	for data := range data_chan {
		_, err := dec.Decode(data, output)
		if err != nil {
			log.Println("cannot decode opus")
		}
		err = streamer.Write()
		if err != nil && err != portaudio.OutputUnderflowed {
			log.Printf("error write to audio stream : %v", err)
			return
		}
		// streamer.Write()
	}

}
func main() {
	data_chan := make(chan []byte, 1000)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go mp3_to_opus("305.mp3", data_chan, wg)
	go opus_to_mp3(data_chan, wg)
	wg.Wait()
	// binary.Read()
	// encoder.Encode()
}
