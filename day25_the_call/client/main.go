package main

import (
	"client/output"
	"client/protobuf"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	net, err := grpc.Dial("192.168.1.9:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	client, _ := protobuf.NewTheCallClient(net).Calling(ctx)
	sig := make(chan struct{})
	out, in := io.Pipe()
	output, _ := output.New(1024, out)
	go func() {
		for {
			data, err := client.Recv()
			in.Write(data.GetSound())
			if err != nil {
				in.Close()
				sig <- struct{}{}
			}
		}
	}()
	go func() {
		output.Play()
	}()
	<-sig
}
