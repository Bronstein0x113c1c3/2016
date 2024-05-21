package main

import (
	pb "client/protobuf"
	"context"
	"fmt"
)

func listAllSong(songslist []*pb.SongInfo) {
	fmt.Printf("%-10v|%-50v|%-50v|%-50v|%-10v|%-4v|\n", "ID", "Name", "Artists", "Album", "Track", "Year")
	for _, j := range songslist {
		fmt.Println(eachof(j))
	}
}
func eachof(info *pb.SongInfo) string {
	// var res string
	// res += fmt.Sprintf("ID: %v \n", info.Info.Id)
	// res += fmt.Sprintf("Name: %v \n", info.Info.Name)
	// res += fmt.Sprintf("Artist: %v \n", info.Info.ArtistName)
	// res += fmt.Sprintf("Album: %v \n", info.Info.AlbumName)
	// res += fmt.Sprintf("Track: %v \n", info.Info.TrackNumber)
	// res += fmt.Sprintf("Year: %v \n", info.Info.PublishedYear)
	return fmt.Sprintf("%-10v|%-50v|%-50v|%-50v|%-10v|%-4v|", info.Id, info.Name, info.ArtistName, info.AlbumName, info.TrackNumber, info.PublishedYear)
}
func GetAll(client pb.PlayerClient) error {
	ctx := context.Background()
	req, err := client.GetAllSong(ctx, &pb.SongRequest{
		Request: &pb.SongRequest_S{
			S: &pb.Signal{},
		},
	})
	if err != nil {
		return err
	}
	listAllSong(req.Songs)
	// PlayMusic(56, client)
	return nil
}
