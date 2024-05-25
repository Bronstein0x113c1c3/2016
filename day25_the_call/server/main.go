package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"server/caller"
	in "server/input"
	pb "server/protobuf"
	"sync"
	"syscall"

	"net/http"
	_ "net/http/pprof"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer stop()
	// io.MultiReader()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	input, err := in.New(1024)
	//if you don't like it, just
	go input.Play()
	// go input.Pause()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("The microphone input is initiated")
	service := caller.New("127.0.0.1", 8080)
	// defer service.Close()
	grpc_helper := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterTheCallServer(grpc_helper, service)
	net, err := net.Listen("tcp", fmt.Sprintf("%v", service))
	log.Println("Service, server is implemented!")
	if err != nil {
		log.Fatalln(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			var s string
			fmt.Print("Press a command (play, pause, stop): ")
			fmt.Scanln(&s)
			select {
			case <-sigs:
				log.Println("Signal after forcing initiating")
				go input.Stop()
				go grpc_helper.GracefulStop()

				go service.Close()
				wg.Done()
				return
			default:
				switch s {
				case "pause":
					fmt.Println("Paused")
					// clear()
					input.Pause()
					continue
				case "play":
					fmt.Println("Continue")
					input.Play()
					continue

					// output.Play()
				case "stop":
					fmt.Println("Prepare to stop all....")
					go input.Stop()
					fmt.Println("Stopped input....")
					grpc_helper.GracefulStop()
					fmt.Println("Stopped RPCs....")
					wg.Done()

					// output.Stop()
					// stop <- struct{}{}

					return
				default:
					fmt.Println("Wrong command!!!")
					continue

				}
			}
		}
	}()
	wg.Add(1)
	go func() {
		stream := input.GetStream()
		for {
			data := make([]byte, 1024)
			n, err := stream.Read(data)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			select {
			case x := <-service.DeleteSignal:
				if x == "Shutting down all!!!" {
					wg.Done()
					return
				} else {
					log.Println("Waiting for ....")
					<-service.WaitSignal
					continue
				}
			default:
				service.Update(data[0:n])
			}
		}
	}()
	go grpc_helper.Serve(net)
	go func() {
		fmt.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()
	wg.Wait()
	return
	//

}
