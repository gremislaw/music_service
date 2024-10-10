package core

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"sync"
)

type SimplePlaylist struct {
	Name string
	Ctx context.Context
	Songs *list.List

	currentSongNode *list.Element
	currentSongPlayTime int
	pauseCtx context.CancelFunc
	coreMtx *sync.Mutex
	isPlaying bool
}

func CreateSimplePlaylist(name string, songs *list.List, ctx context.Context) *SimplePlaylist {
	p := &SimplePlaylist{Name: name, Songs: songs, Ctx: ctx}

	p.coreMtx = new(sync.Mutex)
	p.currentSongPlayTime = 0
	p.isPlaying = false
	p.pauseCtx = nil
	p.currentSongNode = p.Songs.Front()
	return p
}

func AddPlaylist(playlist *SimplePlaylist) string {
	res := fmt.Sprintf("New playlist added. Name: %s.", playlist.Name)
	fmt.Println(res)
	return res
}

func DeletePlaylist(name string, p *SimplePlaylist) (string, error) {
	var err error
	res := fmt.Sprintf("Playlist deleted. Name: %s", p.Name)
	if (p.Name != name) {
		return res, err
	}
	if p.isPlaying {
		playing := "playlist is playing, can't delete"
		err = errors.New(playing)
		return playing, err
	}
	return res, err
}

func GetPlaylist(playlist *SimplePlaylist) string {
	res := fmt.Sprintf("Playlist chosen. Name: %s.", playlist.Name)
	fmt.Println(res)
	return res
}