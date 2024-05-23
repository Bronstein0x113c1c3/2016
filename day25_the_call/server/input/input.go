package in

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/gordonklaus/portaudio"
)

type Input struct {
	// in_stream_channel chan []int16
	close_signal chan struct{}
	// input_stream  io.Writer
	output_stream io.Writer
}

// can i do it with io.Pipe??????? maybe
func (i *Input) GetStream() io.Writer {
	return i.output_stream
}

// func (i *Input) GetChannel() chan []int16 {
// 	return i.in_stream_channel
// }

// func (o *Input) Terminate() error {
// 	// close(o.in_stream_channel)
// 	// o.stream.Stop()
// 	// // close(o.in_stream_channel)
// 	o.close_signal <- struct{}{}
// 	if err := o.stream.Close(); err != nil {
// 		return err
// 	}
// 	if err := portaudio.Terminate(); err != nil {
// 		return err
// 	}
// 	// close(o.in_stream_channel)

// 	// close(o.close_signal)
// 	// // close(sound_chan)
// 	// log.Println("Closed successfully!!!")

// 	return nil
// }
// func (o *Input) Start() {
// 	o.stream.Start()
// }

func (i *Input) Stop() {
	// i.close_signal <- struct{}{}
	close(i.close_signal)
}

func New() (*Input, error) {
	portaudio.Initialize()
	buffer := make([]int16, 8196)
	// data_chan := make(chan []int16)
	var buffer_stream bytes.Buffer
	signal := make(chan struct{})

	stream, err := portaudio.OpenDefaultStream(1, 0, 16000, len(buffer), &buffer)

	if err != nil {
		return nil, err
	}
	stream.Start()
	go func(signal chan struct{}) {
		for {
			stream.Read()
			// x := buffer
			binary.Write(&buffer_stream, binary.LittleEndian, buffer)
			select {
			case <-signal:
				stream.Stop()
				portaudio.Terminate()
				return
			default:

			}
		}
	}(signal)
	return &Input{
		close_signal: signal,
		// input_stream:  in,
		output_stream: &buffer_stream,
	}, nil
	// go func(data_chan chan []int16) {
	// 	data_chan <- in
	// }(data_chan)
	// time.Sleep(time.Millisecond * 10)
}
