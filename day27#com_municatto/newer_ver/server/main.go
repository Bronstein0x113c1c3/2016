package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"serv/interceptor"
	"serv/protobuf"
	"serv/serverimpl"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	defer log.Println("Closed completed!!!")
	input_chan := make(chan serverimpl.Chunk)
	server := serverimpl.New("", 8080, input_chan)
	grpc_helper := grpc.NewServer(grpc.ChainStreamInterceptor(interceptor.Limiting, interceptor.ChannelFinding(server)))

	protobuf.RegisterCallingServer(grpc_helper, server)
	defer close(input_chan)
	defer grpc_helper.GracefulStop()
	log.Printf("Setting up done!!! %v \n", fmt.Sprintf("%v", *server))
	lis, err := net.Listen("tcp", fmt.Sprint(server))
	if err != nil {
		log.Fatalln(err)
		return
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer func() {
			for _, client := range server.Output {
				if client != nil {
					close(client)
				}
			}
		}()
		for {

			select {
			case <-sigs:
				done <- struct{}{}
				return
			case x := <-server.ChangeSignal:
				log.Println("Something changed....")
				if x == 0 {
					log.Println("Added someone.... Resetting")
					continue
				} else {
					i := -x - 1
					log.Printf("%v is requested for deletion \n", i)
					close(server.Output[i])
					server.Output[i] = nil
					continue
				}

			case data, ok := <-input_chan:
				_, _, id := data.Get()
				if !ok {
					log.Println("Channel is forcibly closed")
					for i := range server.Output {
						if server.Output[i] != nil {
							close(server.Output[i])
						}
					}
				}
				for i := range server.Output {
					if server.Output[i] != nil && id != i {
						server.Output[i] <- data
					}
				}
			}
		}
	}()

	go grpc_helper.Serve(lis)
	<-done
	// <-sigs
	// <-sigs
}
