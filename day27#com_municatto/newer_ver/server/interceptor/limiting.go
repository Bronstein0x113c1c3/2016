package interceptor

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var ConnLimit chan struct{}

func init() {
	ConnLimit = make(chan struct{}, 10)
}

// func(srv interface{}, ss ServerStream, info *StreamServerInfo,handler StreamHandler) error
func Limiting(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	p, ok := peer.FromContext(ss.Context())
	log.Printf("%v is called... \n", info.FullMethod)
	if !ok {
		log.Println("Cannot get any context.....")
	}
	log.Printf("%v wants to connect....: ", p.Addr.String())

	a := 10 - len(ConnLimit)
	if a <= 0 {
		// log.Printf("remaining %v connections... \n", a)
		log.Println("Out of slots")
		return status.Error(codes.ResourceExhausted, "out of slots, cancelled!!!")
	}
	log.Println("Accepted!!!")
	ConnLimit <- struct{}{}
	a = 10 - len(ConnLimit)
	log.Printf("remaining %v connections... \n", a)
	handler(srv, ss)
	return nil
}
