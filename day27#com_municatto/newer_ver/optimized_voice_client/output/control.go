package output

import "github.com/gordonklaus/portaudio"

func (o *Output) Close() error {
	defer portaudio.Terminate()
	o.stream.Stop()
	return o.stream.Close()
}
