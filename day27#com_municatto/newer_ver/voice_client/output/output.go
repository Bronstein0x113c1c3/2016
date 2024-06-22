package output

import (
	"encoding/binary"
	"io"

	"github.com/gordonklaus/portaudio"
)

type Output struct {
	input_stream io.Reader
	stop_signal  chan string
}

func (o *Output) Stop() {
	o.stop_signal <- "stop"
}
func (o *Output) Play() {
	o.stop_signal <- "play"
}
func (o *Output) Pause() {
	o.stop_signal <- "pause"
}
func New(buffer_size int, input_stream io.Reader) (*Output, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}
	buf := make([]int16, buffer_size)
	reading, err := portaudio.OpenDefaultStream(0, 2, 16000, len(buf), &buf)
	if err != nil {
		return nil, err
	}
	stop_signal := make(chan string)
	go func() {
		defer input_stream.(*io.PipeReader).Close()
		for {
			_ = binary.Read(input_stream, binary.LittleEndian, &buf)
			reading.Write()
			select {
			case x := <-stop_signal:
				if x == "stop" {
					reading.Close()
					portaudio.Terminate()
					return
				}
				if x == "play" {
					reading.Start()
					continue
				}
				if x == "pause" {
					reading.Abort()
					clear(buf)
					continue
				}
			default:
			}
		}
	}()
	return &Output{
		stop_signal:  stop_signal,
		input_stream: input_stream,
	}, nil
}
