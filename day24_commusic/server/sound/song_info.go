package sound

import (
	"os"

	"github.com/dhowden/tag"
)

type Song struct {
	id             int
	name           string
	artist_name    string
	album_name     string
	track_number   int
	published_year int
	file_name      string
}

func GetInfoSongFLAC(root string, file_name string, id int) Song {
	f, err := os.Open(root + file_name)
	if err != nil {
		// log.Fatal(err)
		return Song{}
	}
	meta, err := tag.ReadFrom(f)
	if err != nil {
		// log.Fatal(err)
		return Song{}
	}
	i, _ := meta.Track()
	// fmt.Println(res)
	return Song{
		id:             id,
		name:           meta.Title(),
		artist_name:    meta.Artist(),
		album_name:     meta.Album(),
		track_number:   i,
		published_year: meta.Year(),
		file_name:      file_name,
	}

}

func GetInfoSongMP3(root string, file_name string, id int) Song {
	f, err := os.Open(root + file_name)
	if err != nil {
		// log.Fatal(err)
		return Song{}
	}
	meta, err := tag.ReadFrom(f)
	if err != nil {
		// log.Fatal(err)
		return Song{}
	}
	i, _ := meta.Track()
	// fmt.Println(res)
	return Song{
		id:             id,
		name:           meta.Title(),
		artist_name:    meta.Artist(),
		album_name:     meta.Album(),
		track_number:   i,
		published_year: meta.Year(),
		file_name:      file_name,
	}

}

func (s Song) GetID() int {
	return s.id
}

func (s Song) GetName() string {
	return s.name
}

func (s Song) GetArtist() string {
	return s.artist_name
}

func (s Song) GetAlbum() string {
	return s.album_name
}

func (s Song) GetTrackNum() int {
	return s.track_number
}
func (s Song) GetPublishedYear() int {
	return s.published_year
}
func (s Song) GetFileName() string {
	return s.file_name
}
