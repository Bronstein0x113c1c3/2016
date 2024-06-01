package in

import (
	"encoding/binary"
	"io"
	"log"

	"github.com/gordonklaus/portaudio"
)

type Input struct {
	// in_stream_channel chan []int16
	sig chan string
	// input_stream  io.Reader
	output_stream io.Reader
}

// can i do it with io.Pipe??????? maybe
func (i *Input) GetStream() io.Reader {
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

func (i *Input) Play() {
	i.sig <- "play"
}
func (i *Input) Stop() {
	// i.close_signal <- struct{}{}
	// close(i.sig)
	log.Println("Stopping called")
	// i.output_stream.(*io.PipeReader).Close()
	i.sig <- "stop"
}
func (i *Input) Pause() {
	// i.pause_signal <- struct{}{}
	i.sig <- "pause"
}

func New(buffer_size int) (*Input, error) {
	portaudio.Initialize()
	// mutex := &sync.Mutex{}
	buffer := make([]int16, buffer_size)
	out, in := io.Pipe()
	signal := make(chan string)
	stream, err := portaudio.OpenDefaultStream(2, 0, 16000, len(buffer), &buffer)
	// stream.Start()
	if err != nil {
		return nil, err
	}
	go func(signal chan string) {
	loop:
		for {
			stream.Read()
			// x := buffer
			binary.Write(in, binary.LittleEndian, buffer)
			clear(buffer)
			select {
			case x := <-signal:

				switch x {
				case "pause":
					// mutex.TryLock()
					stream.Abort()
					clear(buffer)
					// mutex.Unlock()
					goto loop
				case "stop":
					// mutex.Lock()
					stream.Close()
					portaudio.Terminate()
					in.Close()
					out.Close()
					// mutex.Unlock()
					return
				case "play":
					// mutex.Lock()
					log.Println("Play is called")
					stream.Start()
					// mutex.Unlock()
					goto loop
				}
			default:

			}
		}
	}(signal)
	return &Input{
		sig: signal,
		// input_stream:  in,
		output_stream: out,
	}, nil
	// go func(data_chan chan []int16) {
	// 	data_chan <- in
	// }(data_chan)
	// time.Sleep(time.Millisecond * 10)
}
