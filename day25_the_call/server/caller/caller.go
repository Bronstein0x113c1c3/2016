package caller

import "github.com/gordonklaus/portaudio"

type Caller struct {
	host  string
	port  int
	input *portaudio.Stream
	// output *portaudio.Stream
}
