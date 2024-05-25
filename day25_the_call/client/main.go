package main

import (
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
	// ctx := context.Background()
	// ctx.
	go func() {
		for {
			data, err := client.Recv()
			in.Write(data.GetSound())
			if err != nil {
				in.Close()
				output.Stop()
				// sig <- struct{}{}
			}
		}
	}()

	go func() {
		output.Play()
	}()
	// go func() {
	// 	for {
	// 		select {
	// 		case <-sigs:
	// 			client.Send()
	// 		}
	// 	}
	// }()
	<-sigs
	// client.CloseSend()
	client.Send(&protobuf.Client_MSGSound{
		Sound: []byte("Goodbye!!"),
	})

	log.Println("CloseSend signal is sent!!!")
	in.Close()
	net.Close()

}
