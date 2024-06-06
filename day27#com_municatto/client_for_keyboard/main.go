package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	pb "client/protobuf"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.NewClient("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client, err := pb.NewCallingClient(conn).VoIP(context.Background())
	defer client.CloseSend()
	if err != nil {
		log.Fatalln(err)
	}
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, os.Interrupt, os.Kill)
	// // signal.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	ctx2, cancel2 := context.WithCancel(ctx)
	defer cancel2()
	wg := &sync.WaitGroup{}
	var name string
	fmt.Print("Press your name: ")
	fmt.Scanln(&name)
	wg.Add(2)

	type chunk struct {
		name string
		ch   string
	}
	data_chan := make(chan chunk)
	defer close(data_chan)
	go func() {
		for {
			data, err := client.Recv()
			if err != nil {
				// close(data_chan)
				return
			}
			data_chan <- chunk{name: data.Msg.Name, ch: string(data.Msg.Chunk)}
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
				log.Printf("%v: %v \n", x.name, x.ch)
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				log.Println("Context sending side called")
				return
			default:
				scanner.Scan()
				s := scanner.Text()
				client.Send(&pb.ClientMSG{
					Chunk: []byte(s),
					Name:  name,
				})

			}
		}
	}()
	wg.Wait()
	client.CloseSend()
	log.Println("Exited successfully!!!")
}
