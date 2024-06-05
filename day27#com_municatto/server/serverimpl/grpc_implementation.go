package serverimpl

import (
	"log"
	pb "serv/protobuf"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// implement /run/media/jonathan0x113c1c3/New Volume/immortal/2016/day27#com_municatto/protobuf/protobuffor each connection to server
func (s *Server) VoIP(conn pb.Calling_VoIPServer) error {
	// pb.ClientMSG

	log.Println("Knock knock....")
	id, err := s.add()

	log.Printf("Added %v \n", id)
	if err != nil {
		return status.Error(codes.Aborted, "The server is full, cannot add!!!")
	}
	sig1 := make(chan struct{})
	sig2 := make(chan struct{})

	//receive the byte to send to the channel
	go receive(conn, s.Input, id, sig1)
	//get the data then send to the client
	go send(conn, s.Output[id], id, sig2)
	for {
		select {
		case <-sig1:
			log.Println("Disconnected....")
			s.delete_chan(id, false)
			return nil
		case <-sig2:
			log.Println("Forcing closed!!!!")
			s.delete_chan(id, true)
			return nil
		}
	}

}
