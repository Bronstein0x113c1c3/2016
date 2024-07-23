package main

import (
	pb "client/protobuf"
	"crypto/tls"
	"log"

	grpcquic "github.com/coremedic/grpc-quic"
	"google.golang.org/grpc"
)

func main() {
	// Create TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"music-player"},
	}
	creds := grpcquic.NewCredentials(tlsConfig)

	// Connect to gRPC Service Server
	dialer := grpcquic.NewQuicDialer(tlsConfig)
	grpcOpts := []grpc.DialOption{
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial("localhost:8080", grpcOpts...)
	if err != nil {
		log.Fatal(err)
	}

	// Close connection at end of function
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	// Create gRPC client
	grpcClient := pb.NewPlayerClient(conn)

	// Send gRPC request
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// req, err := grpcClient.GetAllSong(ctx, &pb.SongRequest{
	// 	Request: &pb.SongRequest_S{
	// 		S: &pb.Signal{},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("[gRPC]: %v\n", req.Songs)
	GetAll(grpcClient)
	PlayMusic(58, grpcClient)

}
