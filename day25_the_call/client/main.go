package main

import (
	"client/input"
	"client/output"
	"client/protobuf"
	"context"
	"crypto/tls"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	grpcquic "github.com/coremedic/grpc-quic"
	"google.golang.org/grpc"
)

func main() {
	// Create TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"caller"},
	}
	creds := grpcquic.NewCredentials(tlsConfig)

	// Connect to gRPC Service Server
	dialer := grpcquic.NewQuicDialer(tlsConfig)
	grpcOpts := []grpc.DialOption{
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(creds),
	}
	net, err := grpc.Dial("192.168.1.4:8080", grpcOpts...)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, _ := protobuf.NewTheCallClient(net).Calling(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	out, in := io.Pipe()
	output, _ := output.New(1024, out)
	input, _ := input.New(1024)

	go func() {
		for {
			data, err := client.Recv()
			in.Write(data.GetSound())
			if err != nil {
				in.Close()
				output.Stop()
				input.Stop()
				return

			}
		}
	}()
	go func() {
		stream := input.GetStream()
		for {
			data := make([]byte, 4096)
			n, err := stream.Read(data)
			if err != nil {
				log.Println("EOF signal received!!! Stream is stopped!!! Start to closing data channel")

				return
			}
			client.Send(&protobuf.Client_MSGSound{Sound: data[0:n]})

		}

	}()
	go func() {
		output.Play()
		// input.Play()
	}()

	<-sigs
	client.CloseSend()

	log.Println("CloseSend signal is sent!!!")
	in.Close()
	net.Close()

}
