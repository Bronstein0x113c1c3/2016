package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"server/caller"
	in "server/input"
	pb "server/protobuf"
	"server/utility"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// io.MultiReader()
	input, err := in.New(1024)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("The microphone input is initiated")
	service := caller.New("192.168.1.4", 8080)
	grpc_helper := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	pb.RegisterTheCallServer(grpc_helper, service)

	//time for http3 to shine!!!

	// // net, err := net.Listen("tcp", fmt.Sprintf("%v", service))
	tlsconfig, err := utility.GenerateTLSConfig("caller")
	if err != nil {
		log.Fatalln("Failed creating tls config for server")
	}

	// quic_tunnel, err := quic.ListenAddr(fmt.Sprint(service), tlsconfig, nil)
	// if err != nil {
	// 	log.Fatalln("Failed creating quic/http3 tunnel")
	// }
	// //parsing quic listener to universal net listener!!
	// net := grpcquic.Listen(*quic_tunnel)
	net, err := utility.NewHTTP3Lis(fmt.Sprint(service), tlsconfig)
	log.Println("Service, server is implemented, with HTTP/3!")
	// defer grpc_helper.GracefulStop()

	if err != nil {
		log.Fatalln(err)
	}
	go input.Play()
	data_chan := make(chan []byte, 50)
	// change := make(chan struct{})
	go func() {
		for {
			var s string
			log.Println("Press a command (play, pause, stop, force shutdown - ctrl + c)...")
			// fmt.Scanln(&s)
			select {
			case <-sig:
				log.Println("Forcing shutting down signal received!!!")
				log.Println("Stopping input!!!")
				input.Stop()
				log.Println("Stopped input!!!")
				grpc_helper.GracefulStop()
				log.Println("Stopped gRPC!!!")

				return
			default:
				fmt.Scanln(&s)
				switch s {
				case "play":
					input.Play()
					continue
				case "pause":
					input.Pause()
					continue
				case "stop":
					input.Stop()
					input.GetStream().(*io.PipeReader).Close()
					grpc_helper.GracefulStop()
					log.Println("Stopped gRPC!!!")
					return
				}
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
				close(data_chan)
				return
			}
			// log.Println("Exchanging....")
			data_chan <- data[0:n]
		}

	}()
	go func(data_chan chan []byte) {
		defer log.Println("Don't send anymore....")
		for {
			data, ok := <-data_chan
			if !ok {
				log.Println("Stream is closed....")
				for i, client := range service.ListOfClient {
					log.Printf("%v is starting to be closed \n", i)
					close(client)
				}
				// clear()
				return
			}
			select {
			case x := <-service.ChangeSignal:
				log.Println("Something changed....")
				if x == 0 {
					continue
				} else {
					close(service.ListOfClient[x])
					// service.Mutex.Lock()
					delete(service.ListOfClient, x)
					// service.Mutex.Lock()
					// log.Printf("There are %v active connections \n", service.GetAmountOfChannel())
					continue
				}
			default:
				for _, client := range service.ListOfClient {
					client <- data
				}
				// for i := range service.GetAmountOfChannel() {
				// 	if c, ok := caller.ListOfClient.Load(i); ok {
				// 		channel := c.(*chan []byte)
				// 		*channel <- data

				// 	}
				// }
			}
		}
	}(data_chan)

	grpc_helper.Serve(net)
	log.Println("Everything ended!!!")

}

// service.UpdateOrClose(data_chan, service.DeleteSignal)
