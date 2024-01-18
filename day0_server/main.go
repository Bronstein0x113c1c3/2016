package main

import (
	"fmt"
	"net"
)

func handle_conn(c net.Conn) {
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		c.Write([]byte(fmt.Sprintf("%v: Done", c.RemoteAddr())))
	// 		return
	// 	default:
	// 		c.Write([]byte("Bonjour"))
	// 	}
	// }
	fmt.Fprint(c, "Bonjour")
	return
}
func main() {
	// wg := &sync.WaitGroup{}
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer stop()

	listener, _ := net.Listen("tcp", "localhost:1807")
	// go func() {
	// 	<-ctx.Done()
	// }()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle_conn(conn)
	}
}
