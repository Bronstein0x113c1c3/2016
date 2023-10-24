package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var done = make(chan bool, 1)

// func handle_conn(c net.Conn, stop chan os.Signal, wg *sync.WaitGroup) {
// 	for {
// 		select {
// 		case <-stop:
// 			log.Println("Done")
// 			c.Close()
// 			done <- true
// 			wg.Done()
// 			return
// 		default:
// 			io.Copy(os.Stdout, c)

//			}
//		}
//	}
func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(c net.Conn) {
		for {
			select {
			case <-sigs:
				log.Println("Done")
				c.Close()
				done <- true
				wg.Done()
				return
			default:
				var b []byte = make([]byte, 1024)
				c.Read(b)
				fmt.Println(string(b))
				if string(b) == "" {
					log.Println("Done")
					c.Close()
					done <- true
					wg.Done()
				}
			}
		}
	}(conn)

	wg.Wait()
	<-done
}
