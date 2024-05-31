package main

import (
	"client/input"
	"client/output"
	"client/protobuf"
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	net, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
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
			data := make([]byte, 1024)
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
		input.Play()
	}()

	<-sigs
	client.CloseSend()

	log.Println("CloseSend signal is sent!!!")
	in.Close()
	net.Close()

}
