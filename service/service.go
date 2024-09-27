package service

import (
	"container/list"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"music_service/core"
	"time"

	"github.com/gremislaw/music_service/api"
)

type MusicServiceServer struct {
	api.UnimplementedMusicServiceServer
	playlist *core.SimplePlaylist
	db *sql.DB
}

func getSongsFromDb(db *sql.DB) *list.List{
	songs := list.New()
	rows, err := db.Query("select * from playlist")
	if err != nil {
        panic(err)
    }
	defer rows.Close()
	for rows.Next() {
		song := new(core.Song)
		err = rows.Scan(&song.Id, &song.Duration, &song.Name, &song.Author)
		if err != nil {
			panic(err)
		}
		songs.PushBack(song)
	}
	return songs
}

func isDBAvailable(db *sql.DB) bool{
	var res bool = true

	_, err := db.Query("select 1")
	
	if err != nil {
		res = false
    }
	return res
}

func NewService(db *sql.DB) api.MusicServiceServer {
	service := MusicServiceServer{}

	for i := 1; !isDBAvailable(db); i++ {
        fmt.Printf("Db is unavailable(%ds)\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Db is available")
	songs := getSongsFromDb(db)
	playlist := core.CreateSimplePlaylist("My favorite playlist", songs, context.Background())
	service.playlist = playlist
	service.db = db
	return service
}

func (srv MusicServiceServer) Play(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	res := srv.playlist.Play()
	return &api.Response{Response: res}, nil
}

func (srv MusicServiceServer) Pause(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	res := srv.playlist.Pause()
	return &api.Response{Response: res}, nil
}

func (srv MusicServiceServer) AddSong(ctx context.Context, song *api.Song) (*api.Response, error){
	s := new(core.Song)
	s.Author = song.Author
	s.Duration = int(song.Duration)
	s.Name = song.Name
	res := srv.playlist.AddSong(s)
	_, err := srv.db.Exec("insert into playlist (duration, songname, author) values ($1, $2, $3)", song.Duration, song.Name, song.Author)
	if err != nil {
		panic(err)
	}
	return &api.Response{Response: res}, nil
}

func (srv MusicServiceServer) Next(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	res := srv.playlist.Next()
	return &api.Response{Response: res}, nil
}

func (srv MusicServiceServer) Prev(ctx context.Context, empty *api.Empty) (*api.Response, error) {
	res := srv.playlist.Prev()
	return &api.Response{Response: res}, nil
}

func (srv MusicServiceServer) DeleteSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	var res *api.Response = new(api.Response)
	status, err := srv.playlist.DeleteSong(song.Name)
	res.Response = status
	fmt.Println("lolol123")
	if err != nil {
		fmt.Println("Error:", err)
		return res, err
	}
	_, err = srv.db.Exec("delete from playlist where songname = $1", song.Name)
	if err != nil {
		fmt.Println("Error:", err)
		return &api.Response{Response: ""}, err
	}
	return res, err
}

func (srv MusicServiceServer) GetPlaylist(context.Context, *api.Empty) (*api.Playlist, error) {
	list := srv.playlist.GetSongs()
	songsSlice := make([]*api.Song, list.Len())
	i := 0
	var err error = nil

	if list.Len() != 0 {
		for node := list.Front(); node != nil; node = node.Next() {
			songsSlice[i] = &api.Song{
				Author: node.Value.(*core.Song).Author,
				Name: node.Value.(*core.Song).Name,
				Duration: int64(node.Value.(*core.Song).Duration),
			}
			i++
		}
	} else {
		err = errors.New("not found")
	}

	return &api.Playlist{Songs: songsSlice}, err
}

func (srv MusicServiceServer) GetSong(ctx context.Context, song *api.Song) (*api.Song, error) {
	var res *api.Song = nil
	coreSong, err := srv.playlist.GetSong(song.Name)

	if err == nil {
		res = &api.Song{Author: coreSong.Author, Name: coreSong.Name, Duration: int64(coreSong.Duration)}
	}
	return res, err
}

func (srv MusicServiceServer) UpdateSong(ctx context.Context, song *api.Song) (*api.Response, error) {
	err := srv.playlist.UpdateSong(song.Name, song.Author, int(song.Duration))
	response := &api.Response{Response: "song updated"}
	if err != nil {
		return &api.Response{Response: ""}, err
	}
	_, err = srv.db.Exec("update playlist set author = $1, duration = $2 where songname = $3", song.Author, song.Duration, song.Name)
	if err != nil {
		fmt.Println("Error:", err)
		response = &api.Response{Response: ""}
	}
	return response, err
}
