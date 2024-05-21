package server

import (
	"context"
	"fmt"
	"log"
	pb "server/protobuf"
	"server/sound"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) GetAllSong(ctx context.Context, res *pb.SongRequest) (*pb.AllSong, error) {
	log.Println("Called GetAllSong()")
	if _, ok := res.Request.(*pb.SongRequest_Id); ok {
		return nil, status.Error(codes.Aborted, "Bad request!!!")
	}
	// x:= res.Request.(*pb.SongRequest_S)
	list_of_info, err := s.processor.ListAllSong()
	if err != nil {
		return nil, status.Error(codes.Internal, "Error in files")
	}
	if len(list_of_info) <= 0 {
		return nil, status.Error(codes.NotFound, "Songs not found!!!")
	}
	var songlists []*pb.SongInfo
	for _, info := range list_of_info {
		songlists = append(songlists,
			&pb.SongInfo{
				Id:            int32(info.GetID()),
				Name:          info.GetName(),
				ArtistName:    info.GetArtist(),
				AlbumName:     info.GetAlbum(),
				TrackNumber:   int32(info.GetTrackNum()),
				PublishedYear: int32(info.GetPublishedYear()),
			},
		)
	}
	return &pb.AllSong{
		Songs: songlists,
	}, nil
}

func (s *Server) PlaySong(res *pb.SongRequest, srv pb.Player_PlaySongServer) error {
	log.Println("called PlaySong()")
	if _, ok := res.Request.(*pb.SongRequest_S); ok {
		if _, ok := res.Request.(*pb.SongRequest_Id); ok {
			return status.Error(codes.Aborted, "Bad request!!!")
		}
	}
	x := res.Request.(*pb.SongRequest_Id)
	songslist, err := s.processor.ListAllSong()
	if err != nil {
		return status.Error(codes.NotFound, "Songs not found!!!")
	}
	if int(x.Id) < 1 || int(x.Id) > len(songslist) {
		return status.Error(codes.InvalidArgument, "Wrong number, id song must be between 1 and number of the songs!!!")
	}
	// log.Println("dsfsdfdsff")
	song := songslist[x.Id-1]
	log.Println(x.Id)
	first_req := &pb.SongData_Info{
		Info: &pb.SongInfo{
			Id:            int32(song.GetID()),
			Name:          song.GetName(),
			ArtistName:    song.GetArtist(),
			AlbumName:     song.GetAlbum(),
			TrackNumber:   int32(song.GetTrackNum()),
			PublishedYear: int32(song.GetPublishedYear()),
		},
	}
	if err = srv.Send(&pb.SongData{
		Data: first_req,
	}); err != nil {
		return status.Error(codes.Internal, "Implementation failed")
	}
	song_data, err := s.processor.AudioDataStream(song.GetFileName())
	if err != nil {
		return status.Error(codes.Internal, "cannot send the stream")
	}
	for data := range song_data {
		new_chunk := []*pb.Chunk{}
		for _, sample := range data {
			new_chunk = append(new_chunk, &pb.Chunk{
				Left:  sample[0],
				Right: sample[1],
			})
		}
		if err := srv.Send(&pb.SongData{
			Data: &pb.SongData_Chunks{
				Chunks: &pb.Chunks{
					C: new_chunk,
				},
			},
		}); err != nil {
			return err
		}
	}
	log.Printf("Done!!!!")
	return nil

}
