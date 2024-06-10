package interceptor

import (
	"context"
	"log"
	"serv/helper"
	"serv/serverimpl"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// var ConnLimit chan struct{}

// func init() {
// 	ConnLimit = make(chan struct{}, 10)
// }

// func(srv interface{}, ss ServerStream, info *StreamServerInfo,handler StreamHandler) error
type wrapped struct {
	grpc.ServerStream
	id int
}

// for context....

func (w *wrapped) Context() context.Context {
	return context.WithValue(w.ServerStream.Context(), helper.T("channel"), w.id)
}
func ChannelFinding(s *serverimpl.Server) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		id, err := s.Add()
		if err != nil {
			return status.Error(codes.NotFound, "slot not found...")
		}
		log.Printf("slot %v is opened!!!", id)
		handler(srv, &wrapped{ss, id})
		return nil

	}
}
