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
	defer func() {
		wg.Done()
		log.Printf("Closing: %v", conn.RemoteAddr())
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(conn, "Goodbye")
			log.Printf("%v: Goodbye! \n", conn.RemoteAddr())
			return
		default:
			log.Println("Hello")
			fmt.Fprintln(conn, "Hello")
		}
	}
	// fmt.Fprintln(conn, "Hello...")

}
func main() {
	// Set up a TCP listener
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error setting up TCP listener:", err)
		os.Exit(1)
	}
	//step1: setup the context,waitgroup
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	wg := &sync.WaitGroup{}
	defer cancel()
	log.Println("Done setting up")
	/*
		- setup context from signal with the main, listener
		- derive the main context to each smaller connection (maybe uneccessary, but with good impact....)
		- setup to get the signal
	*/
	go func() {
		<-ctx.Done()
	}()
	// Accept connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			break
		}
		wg.Add(1)
		childctx, _ := context.WithCancel(ctx)
		go handleConnection(conn, childctx, wg)
		select {
		case <-ctx.Done():
			log.Println("Starting closing down")
			wg.Wait()
			listener.Close()
			log.Println("Server closed")
			os.Exit(0)
		}
	}
}
