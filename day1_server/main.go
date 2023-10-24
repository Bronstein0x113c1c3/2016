package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
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

func handleConnection(conn net.Conn) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Done!")
			conn.Write([]byte("Goodbye! \n"))
			conn.Close()
			cancel()
			return
		default:
			for {
				conn.Write([]byte("Hello!"))
			}
		}
	}(ctx)

}
func main() {
	// Set up a TCP listener
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error setting up TCP listener:", err)
		os.Exit(1)
	}

	// Accept connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			break
		}

		go handleConnection(conn)
	}
}
