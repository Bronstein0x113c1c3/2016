package server

import (
	"context"
	"fmt"
	pb "server/protobuf"
	"server/sound"
)

type Server struct {
	host string
	port int
	//this is the most important part!!!
	processor *sound.Audio
	//for streaming!!!
	pb.UnimplementedPlayerServer
}

func NewServer(host string, port int, processor *sound.Audio) *Server {
	return &Server{
		host:      host,
		port:      port,
		processor: processor,
	}
}
func (s *Server) String() string {
	return fmt.Sprintf("%v:%v", s.host, s.port)
}

func (s *Server) GetAllSong(ctx context.Context, res *pb.SongRequest) (*pb.SongInfo, error) {
	return nil, nil
}

func (s *Server) PlaySong(res *pb.SongRequest, srv pb.Player_PlaySongServer) (*pb.SongInfo, error) {
	return nil, nil
}
