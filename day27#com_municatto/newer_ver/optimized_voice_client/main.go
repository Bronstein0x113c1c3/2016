package main

import (
	"client/input"
	"client/output"
	pb "client/protobuf"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

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
func main() {
	log.Println("connecting to the server....")
	client, err := init_the_client("localhost:8080", "caller")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.CloseSend()
	wg := &sync.WaitGroup{}

	log.Println("connected, init the i/o...")
	input, err := input.InputInit(DefaultChannels, DefaultSampleRate, DefaultFrameSize, DefaultOpusDataLength, wg)
	if err != nil {
		return
	}
	data_chan := make(chan []byte, 1000)
	// defer close(data_chan)
	output, err := output.OutputInit(DefaultChannels, DefaultSampleRate, DefaultFrameSize, data_chan)
	if err != nil {
		return
	}
	log.Println("io done, initiating the signal....")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	wg.Add(1)
	log.Println("signal done, start processing....")

	log.Print("Your name?: ")
	var name string
	fmt.Scanln(&name)

	go input.Process()
	go func() {
		data_chan := input.GetChannel()
		for data := range data_chan {
			client.Send(&pb.ClientMSG{
				Chunk: data,
				Name:  name,
			})
		}
		// for data:= range
	}()
	go func() {
		// data, err := client.Recv()
		// if err != nil {
		// 	stop()
		// 	return
		// }
		// data_chan <- data.Msg.Chunk
		for {
			select {
			case _, ok := <-ctx.Done():
				if !ok {
					return
				}
			default:
				data, err := client.Recv()
				if err != nil {
					stop()
					return
				}
				data_chan <- data.Msg.Chunk
			}
		}
	}()
	go output.Process()
	<-ctx.Done()
	go input.Close()
	go output.Close()
	defer wg.Wait()
}
