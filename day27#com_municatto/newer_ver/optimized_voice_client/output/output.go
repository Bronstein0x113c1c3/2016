package output

import (
	"github.com/gordonklaus/portaudio"
	// "github.com/hraban/opus"
	"gopkg.in/hraban/opus.v2"
)

// opus decoder....

type Output struct {
	buf     []int16
	stream  *portaudio.Stream
	decoder *opus.Decoder
	// wg        *sync.WaitGroup
	data_chan chan []byte
}

func OutputInit(channel int, sample_rate float32, buffer_size int, data_chan chan []byte) (*Output, error) {
	buf := make([]int16, buffer_size)
	portaudio.Initialize()
	streamer, err := portaudio.OpenDefaultStream(0, channel, float64(sample_rate), buffer_size, &buf)
	if err != nil {
		portaudio.Terminate()
		return nil, err
	}
	decoder, err := opus.NewDecoder(int(sample_rate), channel)
	if err != nil {
		portaudio.Terminate()
		return nil, err
	}
	return &Output{
		buf:       buf,
		stream:    streamer,
		decoder:   decoder,
		data_chan: data_chan,
	}, nil
}

func (o *Output) Process() {
	// defer o.wg.Done()
	o.stream.Start()
	for data := range o.data_chan {
		o.decoder.Decode(data, o.buf)
		o.stream.Write()
	}
}

