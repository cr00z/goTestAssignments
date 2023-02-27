package playlist

import (
	"container/list"
	"gocloudcamp/internal/domain/song"
	"sync"
	"time"
)

type Playlist struct {
	sync.Mutex
	items          *list.List
	currentSong    *list.Element
	currentSongPos time.Duration
	played         bool
}

func NewPlaylist() *Playlist {
	return &Playlist{
		items: list.New(),
	}
}

func (pl *Playlist) CurrentSong() (*song.Song, error) {
	song, ok := pl.currentSong.Value.(*song.Song)

	if !ok {
		return nil, ErrorSongCastError
	}

	return song, nil
}

func (pl *Playlist) CurrentSongPos() time.Duration {
	return pl.currentSongPos
}

func (pl *Playlist) SetCurrentSongPos(currentSongPos time.Duration) {
	pl.currentSongPos = currentSongPos
}

func (pl *Playlist) Played() bool {
	return pl.played
}

func (pl *Playlist) Empty() bool {
	return pl.items.Len() == 0
}
func (pl *Playlist) Play() error {
	if pl.Empty() {
		return ErrorPlaylistIsEmpty
	}

	if pl.played {
		return ErrorPlaylistIsAlreadyPlaying
	}
	pl.played = true

	return nil
}

func (pl *Playlist) Pause() error {
	if !pl.played {
		return ErrorPlaylistIsAlreadyPaused
	}
	pl.played = false

	return nil
}

func (pl *Playlist) AddSong(song *song.Song) {
	pl.Lock()
	defer pl.Unlock()

	pl.items.PushBack(song)
	if pl.items.Len() == 1 {
		pl.currentSong = pl.items.Front()
	}
}

// Next plays the next song. If the playlist is over, skips to the beginning
func (pl *Playlist) Next() error {
	pl.Lock()
	defer pl.Unlock()

	if pl.Empty() {
		return ErrorPlaylistIsEmpty
	}

	pl.currentSong = pl.currentSong.Next()
	if pl.currentSong == nil {
		pl.currentSong = pl.items.Front()
	}

	pl.currentSongPos = 0

	return nil
}

// Prev plays the previous song. If the playlist is over, skips to the end
func (pl *Playlist) Prev() error {
	pl.Lock()
	defer pl.Unlock()

	if pl.Empty() {
		return ErrorPlaylistIsEmpty
	}

	pl.currentSong = pl.currentSong.Prev()
	if pl.currentSong == nil {
		pl.currentSong = pl.items.Back()
	}

	pl.currentSongPos = 0

	return nil
}

func (pl *Playlist) First() error {
	pl.Lock()
	defer pl.Unlock()

	if pl.Empty() {
		return ErrorPlaylistIsEmpty
	}

	pl.currentSong = pl.items.Front()
	return nil
}
