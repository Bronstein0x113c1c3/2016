package in

import (
	"encoding/binary"
	"io"
	"log"
	"sync"

	"github.com/gordonklaus/portaudio"
)

type Input struct {
	// in_stream_channel chan []int16
	sig chan string
	// input_stream  io.Reader
	output_stream io.Reader
	// input_stream  io.Writer
	mu *sync.RWMutex
	// stream        *portaudio.Stream
}

// can i do it with io.Pipe??????? maybe
func (i *Input) GetStream() io.Reader {
	return i.output_stream
}

func (i *Input) Play() {
	log.Println("play is called")
	i.sig <- "play"
	// i.mu.RLock()
	// i.stream.Start()
	// i.mu.Unlock()
}
func (i *Input) Stop() {

	// i.mu.Lock()
	// i.stream.Close()
	// i.input_stream.(*io.PipeWriter).Close()
	// i.mu.Unlock()
	// i.close_signal <- struct{}{}
	// close(i.sig)
	log.Println("Stopping called")
	// i.output_stream.(*io.PipeReader).Close()
	i.sig <- "stop"
}
func (i *Input) Pause() {
	log.Println("paused is called")
	// i.mu.RLock()
	// i.stream.Stop()
	// i.input_stream.(*io.PipeWriter).Close()
	// i.mu.Unlock()
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
	// stream.Start()
	go func(stream *portaudio.Stream) {
		for {
			stream.Read()
			err := binary.Write(in, binary.LittleEndian, buffer)
			if err != nil {
				in.Close()
				out.Close()
			}
			select {
			case x := <-signal:
				switch x {
				case "play":
					log.Println("Starting to play!!!")
					stream.Start()
					continue
				case "pause":
					log.Println("Starting to pause!!!")
					clear(buffer)
					stream.Stop()
					continue
				case "stop":
					log.Println("Starting to stop all!!!")
					stream.Close()
					in.Close()
					out.Close()
					close(signal)
					portaudio.Terminate()
					return
				}
			default:
			}
		}
	}(stream)
	// go func(signal chan string) {
	// loop:
	// 	for {
	// 		stream.Read()
	// 		// x := buffer
	// 		binary.Write(in, binary.LittleEndian, buffer)

	// 		select {
	// 		case x := <-signal:

	// 			switch x {
	// 			case "pause":
	// 				// mutex.TryLock()
	// 				stream.Abort()
	// 				clear(buffer)
	// 				// mutex.Unlock()
	// 				goto loop
	// 			case "stop":
	// 				// mutex.Lock()
	// 				stream.Close()
	// 				portaudio.Terminate()
	// 				in.Close()
	// 				// mutex.Unlock()
	// 				return
	// 			case "play":
	// 				// mutex.Lock()
	// 				log.Println("Play is called")
	// 				stream.Start()
	// 				// mutex.Unlock()
	// 				goto loop
	// 			}
	// 		default:

	// 		}
	// 	}
	// }(signal)
	return &Input{
		sig: signal,
		mu:  &sync.RWMutex{},
		// input_stream:  in,
		output_stream: out,
		// stream:        stream,
		// input_stream:  in,
	}, nil
	// go func(data_chan chan []int16) {
	// 	data_chan <- in
	// }(data_chan)
	// time.Sleep(time.Millisecond * 10)
}
