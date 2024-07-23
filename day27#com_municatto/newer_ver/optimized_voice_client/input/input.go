package input

import (
	"sync"

	"github.com/gordonklaus/portaudio"
	"gopkg.in/hraban/opus.v2"
)

// opus encoder....
type Input struct {
	stream    *portaudio.Stream
	buf       []int16
	encoder   *opus.Encoder
	wg        *sync.WaitGroup
	byte_len  int
	signal    chan struct{}
	data_chan chan []byte
}

func InputInit(channel int, sample_rate float32, buffer_size int, data_length int, wg *sync.WaitGroup) (*Input, error) {
	portaudio.Initialize()
	buf := make([]int16, buffer_size)
	streamer, err := portaudio.OpenDefaultStream(channel, 0, float64(sample_rate), buffer_size, &buf)
	if err != nil {
		portaudio.Terminate()
		return nil, err
	}
	encoder, err := opus.NewEncoder(int(sample_rate), channel, opus.AppVoIP)
	if err != nil {
		portaudio.Terminate()
		return nil, err
	}
	// len := int(int(sample_rate) / channel / 1000/)
	data_chan := make(chan []byte, data_length)
	return &Input{
		stream:    streamer,
		encoder:   encoder,
		wg:        wg,
		buf:       buf,
		signal:    make(chan struct{}),
		data_chan: data_chan,
		byte_len:  data_length,
	}, nil
}

func (i *Input) Process() {
	i.stream.Start()
	defer i.wg.Done()
	// defer i.close(i.data_chan)
	defer close(i.data_chan)
	defer portaudio.Terminate()
	defer i.stream.Stop()
	for {
		i.stream.Read()
		data := make([]byte, i.byte_len)
		n, err := i.encoder.Encode(i.buf, data)
		if err != nil {
			return
		}
		select {
		case <-i.signal:
			return
		case i.data_chan <- data[:n]:
		}
	}
}
func (i *Input) GetChannel() chan []byte {
	return i.data_chan
}
