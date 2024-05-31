package caller

import (
	"io"
	"log"
	"server/output"
	pb "server/protobuf"
)

func send_to_client(index uint32, signal chan struct{}, data_chan chan []byte, caller pb.TheCall_CallingServer) {
	defer close(signal)
	// defer s.delete_chan(index, true)
	// defer log.Printf("%v is closed \n", index)
	// defer wg.Done()

	for {
		data, ok := <-data_chan
		if !ok {
			log.Printf("%v is forcing closed \n", index)
			return
		}
		caller.Send(&pb.Server_MSGSound{Sound: data})
	}
}
func hear_from_client(caller pb.TheCall_CallingServer, signal chan struct{}) {
	defer close(signal)
	// internal_sig := make(chan struct{})

	out, in := io.Pipe()
	output, _ := output.New(1024, out, 44100)
	// ctx := context.Background()
	// ctx.
	go func() {
		output.Play()
	}()

	for {
		data, err := caller.Recv()
		in.Write(data.GetSound())
		if err != nil {
			in.Close()
			output.Stop()
			return
			// sig <- struct{}{}
		}
	}
}
