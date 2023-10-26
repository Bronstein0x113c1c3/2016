package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

// func handleConnection(conn net.Conn) {
// 	// Create a channel to listen for OS signals
// 	sigs := make(chan os.Signal, 1)

// 	// Register the channel to receive SIGINT and SIGTERM signals
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	go func() {
// 		// Wait for a signal
// 		sig := <-sigs

// 		fmt.Println()
// 		fmt.Println("Received signal:", sig)

// 		// Close the connection
// 		conn.Close()

// 		fmt.Println("Connection closed")
// 	}()

// 	// Handle the connection here
// 	for {
// 		conn.Write([]byte("Bonjour!"))
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
// 	defer cancel()
// 	go func(ctx context.Context) {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Done!")
// 			conn.Write([]byte("Goodbye! \n"))
// 			conn.Close()
// 			cancel()
// 			return
// 		default:
// 			for {
// 				conn.Write([]byte("Hello!"))
// 			}
// 		}
// 	}(ctx)

// }
func handleConnection(conn net.Conn, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(conn, "Goodbye")
			log.Printf("%v: Goodbye! \n", conn.RemoteAddr())
			conn.Close()
			wg.Done()
			return
		default:
			log.Println("Hello")
			fmt.Fprintln(conn, "Hello")
		}
	}
	// fmt.Fprintln(conn, "Hello...")

}

// func main() {
// 	// Set up a TCP listener
// 	listener, err := net.Listen("tcp", ":3000")
// 	if err != nil {
// 		fmt.Println("Error setting up TCP listener:", err)
// 		os.Exit(1)
// 	}
// 	//step1: setup the context,waitgroup
// 	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
// 	wg := &sync.WaitGroup{}
// 	defer cancel()
// 	conn_chan := make(chan net.Conn)
// 	signal := make(chan bool)
// 	/*
// 		- setup context from signal with the main, listener
// 		- derive the main context to each smaller connection (maybe uneccessary, but with good impact....)
// 		- setup to get the signal
// 	*/

// 	// Accept connections in a loop
// 	wg.Add(1)
// 	go func(wg *sync.WaitGroup, ctx context.Context, s chan bool) {
// 		for {
// 			select {
// 			case <-s:
// 				wg.Done()
// 				return
// 			default:
// 				conn, err := listener.Accept()
// 				if err != nil {
// 					fmt.Println("Error accepting connection:", err)
// 					break
// 				}
// 				conn_chan <- conn
// 			}
// 		}

// 	}(wg, ctx, signal)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			signal <- true
// 			goto ending
// 		case conn := <-conn_chan:
// 			wg.Add(1)
// 			go handleConnection(conn, ctx, wg)
// 		}
// 	}
// ending:
// 	log.Println("Waiting for ....")
// 	wg.Wait()
// 	listener.Close()
// 	log.Println("closed listener...")
// 	os.Exit(0)

// }

//another approach....

func listening(listener net.Listener) <-chan net.Conn {
	res := make(chan net.Conn)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error when accepting...: ", err)
			}
			res <- conn
		}
	}()
	return res
}

func main() {
	defer log.Println("Closed listener....")
	listener, err := net.Listen("tcp", ":3000")
	defer listener.Close()
	if err != nil {
		log.Fatalln("Error when creating listener: ", err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	wg := &sync.WaitGroup{}
	conns := listening(listener)
	defer cancel()
	log.Println("Done preparing....")
	for {
		select {
		case conn := <-conns:
			wg.Add(1)
			go handleConnection(conn, ctx, wg)
		case <-ctx.Done():
			log.Println("Waiting for ....")
			wg.Wait()
			log.Println("closing listener...")
			return
		}
	}
}
