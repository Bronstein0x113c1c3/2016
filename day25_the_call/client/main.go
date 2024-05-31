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

	"github.com/kvark128/minimp3"
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
				return
				// sig <- struct{}{}
			}
		}
	}()
	go func() {
		file, _ := os.Open("rightforyou.mp3")
		decoder := minimp3.NewDecoder(file)
		for {
			buff := make([]byte, 1024)
			n, err := decoder.Read(buff)
			if err != nil {
				return
			}
			// .Write(buff[0:n])
			client.Send(&protobuf.Client_MSGSound{
				Sound: buff[0:n],
			})
		}
		// portaudio.Initialize()
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
	client.CloseSend()
	// client.Send(&protobuf.Client_MSGSound{
	// 	Sound: []byte("Goodbye!!"),
	// })

	log.Println("CloseSend signal is sent!!!")
	in.Close()
	net.Close()

}
