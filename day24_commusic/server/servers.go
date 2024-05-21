package main

import (
	"fmt"
	"log"
	"server/server"
	"server/sound"
	"server/utility"

	pb "server/protobuf"

	grpcquic "github.com/coremedic/grpc-quic"
	"github.com/quic-go/quic-go"
	"google.golang.org/grpc"
)

func main() {
	// testingzone.Readinginfo("./testingzone/11 - Why.flac")

	log.Println("setting up audio files, server")
	audio := sound.NewAudioInstance("./songslist/")
	server := server.NewServer("", 8080, audio)
	grpc_helper := grpc.NewServer()
	pb.RegisterPlayerServer(grpc_helper, server)
	log.Println("done linking to gRPC helper!!")
	log.Println("setting up tunnel")
	tlsconfig, err := utility.GenerateTLSConfig("music-player")
	if err != nil {
		log.Fatalln("Failed creating tls config for server")
	}
	//setup quic tunnel...
	quic_tunnel, err := quic.ListenAddr(fmt.Sprint(server), tlsconfig, nil)
	if err != nil {
		log.Fatalln("Failed creating quic/http3 tunnel")
	}
	//parsing quic listener to universal net listener!!
	net_listener := grpcquic.Listen(*quic_tunnel)
	log.Printf("gRPC-QUIC: listening at %v\n", net_listener.Addr())
	//time for serving
	if err = grpc_helper.Serve(net_listener); err != nil {
		log.Fatalf("failed when listening: %v \n", err)
	}
}
