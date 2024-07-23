package main

import (
	"context"
	"log"
	"net"
	pb "serv/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	helper := grpc.NewServer(grpc.ChainStreamInterceptor(TestContext))
	s := &Serv{}
	pb.RegisterAppServer(helper, s)
	reflection.Register(helper)
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	helper.Serve(lis)
}

type Serv struct {
	pb.UnimplementedAppServer
}

func (s *Serv) Sending(stream pb.App_SendingServer) error {
	log.Println(stream.Context().Value(t("something")))
	for {
		err := stream.Send(&pb.Response{
			X: 101,
		})
		if err != nil {
			return err
		}
	}
}

type Wrapped struct {
	grpc.ServerStream
	i int
}
type t string

func (w *Wrapped) Context() context.Context {
	return context.WithValue(w.ServerStream.Context(), t("something"), int(w.i))
}

func TestContext(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	handler(srv, &Wrapped{ss, 5})
	return nil
}
