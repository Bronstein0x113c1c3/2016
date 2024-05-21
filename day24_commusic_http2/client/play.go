package main

// com-music
import (
	pb "client/protobuf"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func processing_msg(chunks []*pb.Chunk) [][2]float64 {
	var chunk_segment [][2]float64
	for _, c := range chunks {
		var ch [2]float64
		ch[0] = c.Left
		ch[1] = c.Right
		chunk_segment = append(chunk_segment, ch)
	}
	return chunk_segment
}
func receiver(streamer pb.Player_PlaySongClient) chan [][2]float64 {
	sound_stream := make(chan [][2]float64)
	go func() {
		for {
			res, err := streamer.Recv()
			// log.Println(err)
			switch err {
			case io.EOF:
				close(sound_stream)
				log.Println("Done receiving....")
				return
			case nil:
				// log.Println("Nothing error when receiving!!!")
			default:
				close(sound_stream)
				log.Fatalln(err)
				return
			}
			switch x := res.Data.(type) {
			case *pb.SongData_Info:
				id := x.Info.Id
				song_name := x.Info.Name
				artist := x.Info.ArtistName
				track := x.Info.TrackNumber
				album := x.Info.AlbumName
				year := x.Info.PublishedYear
				fmt.Printf("Id: %v \nName: %v \nArtist: %v \nAlbum: %v \nTrack: %v \nYear: %v \n", id, song_name, artist, album, track, year)
			case *pb.SongData_Chunks:
				sound_stream <- processing_msg(x.Chunks.GetC())
				// log.Println("Receving chunk...")
			}
		}
	}()
	return sound_stream
}
func player(sound_chan chan [][2]float64) error {
	done := make(chan struct{})
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second*2))
	streamer := newStreamer(sound_chan, done)
	new_s := beep.Resample(20, sr, beep.SampleRate(40000), streamer)
	speaker.Play(new_s)
	<-done
	return nil
}
func PlayMusic(n int, conn pb.PlayerClient) {
	ctx := context.Background()
	streamer, err := conn.PlaySong(ctx, &pb.SongRequest{
		Request: &pb.SongRequest_Id{
			Id: int32(n),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	streamer.CloseSend()
	data_chan := receiver(streamer)
	player(data_chan)

}
func newStreamer(buffer_chan chan [][2]float64, signal chan struct{}) beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		buffer := <-buffer_chan
		if len(buffer) == 0 {
			signal <- struct{}{}
			return 0, false

		} else {
			for i := range samples {
				samples[i][0] = buffer[i][0]
				samples[i][1] = buffer[i][1]
			}
			return len(samples), true
		}
	})
}
