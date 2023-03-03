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

func (pl *Playlist) checkSongChange(id uint64) error {
	currentSong, err := pl.CurrentSong()
	if err == nil && pl.Played() && currentSong.Id == id {
		return ErrorSongСantBeChanged
	}
	return nil
}

func (pl *Playlist) modifySongFn(modSong *song.Song, fn func(e *list.Element)) error {
	err := pl.checkSongChange(modSong.Id)
	if err != nil {
		return err
	}

	pl.Lock()
	defer pl.Unlock()

	for e := pl.items.Front(); e != nil; e = e.Next() {
		currSong, ok := e.Value.(*song.Song)
		if !ok {
			return ErrorSongCastError
		}

		if currSong.Id == modSong.Id {
			fn(e)
			return nil
		}
	}

	return ErrorSongNotFound
}

func (pl *Playlist) ModifySong(modSong *song.Song) error {
	return pl.modifySongFn(modSong, func(e *list.Element) {
		currSong, _ := e.Value.(*song.Song)
		if modSong.Name != "" {
			currSong.Name = modSong.Name
		}
		if modSong.Duration != 0 {
			currSong.Duration = modSong.Duration
		}
	})
}

func (pl *Playlist) RemoveSong(id uint64) error {
	return pl.modifySongFn(&song.Song{Id: id}, func(e *list.Element) {
		pl.items.Remove(e)
	})
}
