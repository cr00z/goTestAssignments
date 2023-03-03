package player

import (
	"fmt"
	"gocloudcamp/internal/domain/song"
	"time"
)

const SongTick = time.Millisecond * 100

type Player interface {
	CurrentSong() (*song.Song, error)
	CurrentSongPos() time.Duration
	SetCurrentSongPos(currentSongPos time.Duration)
	Played() bool

	Play() error
	Pause() error
	AddSong(song *song.Song)
	Next() error
	Prev() error

	// additional functionality for service
	ModifySong(song *song.Song) error
	RemoveSong(id uint64) error
}

func Start(pl Player) {
	for {
		if pl.Played() {
			song, _ := pl.CurrentSong()
			duration := song.Duration
			fmt.Print(string(song.Name[0]))

			currentSongPos := pl.CurrentSongPos()
			if currentSongPos+SongTick > duration {
				time.Sleep(duration - currentSongPos)
				_ = pl.Next()
			} else {
				time.Sleep(SongTick)
				pl.SetCurrentSongPos(currentSongPos + SongTick)
			}
		}
	}
}
