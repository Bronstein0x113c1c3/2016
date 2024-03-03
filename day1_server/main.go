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

// var stoped = make(chan bool)

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
func handleConnection(conn net.Conn, signal chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-signal:
			fmt.Fprintln(conn, "Goodbye")
			log.Printf("%v: Goodbye! \n", conn.RemoteAddr())
			conn.Close()
			wg.Done()
			return
		default:
			log.Printf("Hello from %v \n", conn.RemoteAddr())
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

func listening(listener net.Listener, signal chan struct{}, wg *sync.WaitGroup) <-chan net.Conn {
	res := make(chan net.Conn)
	// go func() {
	// 	<-ctx.Done()
	// 	log.Println("The listening channel is done!")
	// 	close(res)
	// }()
	// wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer log.Println("Closed the receiving channel.....")
		defer wg.Done()
		defer close(res)
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("The listener is completely down")
				return
			}
			select {
			case <-signal:
				log.Println("Don't accept any connections at that time.....")
				return
			case res <- conn:
			}

			// res <- conn
		}
	}(wg)
	// wg.Wait()
	return res
}

// func listening(listener net.Listener, ctx context.Context) <-chan net.Conn {
// 	res := make(chan net.Conn)
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				close(res)
// 				return
// 			default:
// 				conn, err := listener.Accept()
// 				if err != nil {
// 					fmt.Println("Error when accepting...: ", err)
// 					continue
// 				}
// 				res <- conn
// 			}
// 		}
// 	}()
// 	return res
// }

func main() {
	defer log.Println("Done closing all.....")
	defer log.Println("Closed the connection receiving channel")
	listener, err := net.Listen("tcp", ":3000")
	defer listener.Close()
	if err != nil {
		log.Fatalln("Could not create the listener")
	}

	// done := make(chan struct{})
	//waiting for os.interrupt....
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	// defer func() {
	// 	<-done
	// 	log.Println("The listening receiver is closed.....")
	// }()
	//Done creating the listener....., other preparation.....
	connections_signal := make(chan struct{})
	listener_signal := make(chan struct{})
	// signal_string := make(chan string)
	// defer func(signal_string chan string) {
	// 	<-signal_string
	// }(signal_string)
	//also, the wait group to wait for all the connections to close...
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	/*two channel are used as switch to turn off .....*/
	// connections_receiver channel is used to get the connection after accepting, to make a pipeline....
	wg.Add(1)
	connections_receiver := listening(listener, listener_signal, wg)
	log.Println("Done the preparation......")

	/*the main part....
	- use the loop with:
		- if get os.interrupt signal, push the connections_signal and listener_signal to the listening goroutine and connection goroutines to start closing by closing the channel.
			(alerting mechanism)
		- with each connection from the connection receiver, handle each.
	*/
	defer listener.Close()
	for {
		select {
		case <-ctx.Done():
			// wg.Add(1)
			log.Println("os.Interrupt received from the keyboard")
			close(listener_signal)
			close(connections_signal)
			return
		case conn := <-connections_receiver:
			wg.Add(1)
			go handleConnection(conn, connections_signal, wg)
		}
	}

}
