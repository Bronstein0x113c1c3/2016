package main

import (
	"fmt"
	"log"
	"net"
	"server/server"
	"server/sound"

	pb "server/protobuf"

	"google.golang.org/grpc"
)

func main() {
	// testingzone.Readinginfo("./testingzone/11 - Why.flac")
	// tlsconfig, err := utility.GenerateTLSConfig("music-player")
	log.Println("setting up audio files, server")
	audio := sound.NewAudioInstance("./songslist/")
	server := server.NewServer("0.0.0.0", 8080, audio)
	// var opts []grpc.ServerOption
	// opts = append(opts, grpc.Creds(credentials.NewTLS(tlsconfig)))
	grpc_helper := grpc.NewServer()
	pb.RegisterPlayerServer(grpc_helper, server)
	log.Println("done linking to gRPC helper!!")
	log.Println("setting up tunnel")

	// //setup quic tunnel...
	// quic_tunnel, err := quic.ListenAddr(fmt.Sprint(server), tlsconfig, nil)
	// if err != nil {
	// 	log.Fatalln("Failed creating quic/http3 tunnel")
	// }
	// //parsing quic listener to universal net listener!!
	// net_listener := grpcquic.Listen(*quic_tunnel)
	// log.Printf("gRPC-QUIC: listening at %v\n", net_listener.Addr())
	// //time for serving
	// if err = grpc_helper.Serve(net_listener); err != nil {
	// 	log.Fatalf("failed when listening: %v \n", err)
	// }
	//http2 version
	net, _ := net.Listen("tcp", fmt.Sprint(server))
	if err := grpc_helper.Serve(net); err != nil {
		log.Fatalf("failed when listening: %v \n", err)
	}
}
