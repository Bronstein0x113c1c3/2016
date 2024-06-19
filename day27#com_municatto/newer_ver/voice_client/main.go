package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"

	"client/input"
	"client/output"
	pb "client/protobuf"

	grpcquic "github.com/coremedic/grpc-quic"
	"google.golang.org/grpc"
)

func init_the_client(host string, passcode string) (pb.Calling_VoIPClient, error) {
	passcodes := []string{}
	passcodes = append(passcodes, passcode)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         passcodes,
	}
	creds := grpcquic.NewCredentials(tlsConfig)

	// Connect to gRPC Service Server
	dialer := grpcquic.NewQuicDialer(tlsConfig)
	// grpc
	grpcOpts := []grpc.DialOption{
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.NewClient(host, grpcOpts...)
	if err != nil {
		return nil, err
	}
	client, err := pb.NewCallingClient(conn).VoIP(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

/*
	Paradigm: single input, many outputs from many sources!!!
	Work need to be done:
		- create input, output instances....
		- input -> send to conn...
		- get data -> output -> sound.....
*/

func main() {
	//init the client...
	log.Println("connecting to the server....")
	client, err := init_the_client("localhost:8080", "caller")
	if err != nil {
		log.Fatalln(err)
	}
	// defer client.CloseSend()
	log.Println("connected, init the i/o...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	ctx2, cancel2 := context.WithCancel(ctx)
	defer cancel2()

	//set the input
	input, err := input.New(1024)
	if err != nil {
		log.Println("input error!!!")
	}

	go input.Play()
	defer func() {
		go input.Stop()
	}()
	// defer input.Stop()
	//set the output
	type voice struct {
		output_voice *output.Output
		receiver     io.Writer
	}
	log.Println("connected, init the i/o...")
	output_list := [10]*voice{}

	for i := range 10 {
		// output_list = append(output_list, output.New(1024))
		out, in := io.Pipe()
		output_voice, err := output.New(1024, out)
		if err != nil {
			log.Fatalf("voice output init error!!: %v \n", err)
			return
		}
		voice := &voice{
			output_voice: output_voice,
			receiver:     in,
		}
		go output_voice.Play()
		output_list[i] = voice
	}
	defer func() {
		log.Println("output closing called")
		for i := range 10 {
			output_list[i].receiver.(*io.PipeWriter).Close()
			go output_list[i].output_voice.Stop()
		}
		log.Println("output closing done!!!")

	}()
	log.Println("io init done!!!")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var name string
	fmt.Print("Press your name: ")
	fmt.Scanln(&name)

	//receiving side....
	type chunk struct {
		id   int
		name string
		ch   []byte
	}
	data_chan := make(chan chunk)
	// defer func(data_chan chan chunk) {
	// 	if _, ok := <-data_chan; ok {
	// 		close(data_chan)
	// 	}
	// }(data_chan)
	go func() {
		for {
			data, err := client.Recv()
			if err != nil {
				// close(data_chan)
				log.Println("wait.....")
				// stop()
				// close
				close(data_chan)
				return
			}
			data_chan <- chunk{name: data.Msg.Name, ch: data.Msg.Chunk, id: int(data.Id)}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx1.Done():
				log.Println("Context receving side called")
				return
			case x, ok := <-data_chan:
				if !ok {
					return
				}
				chan_id := x.id
				output_list[chan_id].receiver.Write(x.ch)
			}
		}
	}()

	//sending side....
	go func() {
		stream := input.GetStream()
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				log.Println("Context sending side called")
				return
			default:
				data := make([]byte, 8192)
				n, err := stream.Read(data)
				if err != nil {
					log.Println("EOF signal received!!! Stream is stopped!!! Start to closing data channel")
					return
				}
				client.Send(&pb.ClientMSG{
					Chunk: data[0:n],
					Name:  name,
				})

			}
		}
	}()
	wg.Wait()
	log.Println("closed sending/receiving side....")
	client.CloseSend()
	// input.Stop()
}
