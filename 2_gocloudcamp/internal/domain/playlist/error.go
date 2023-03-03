package playlist

import "errors"

var (
	ErrorSongCastError            = errors.New("song cast error")
	ErrorPlaylistIsAlreadyPlaying = errors.New("playlist is already playing")
	ErrorPlaylistIsAlreadyPaused  = errors.New("playlist is already paused")
	ErrorPlaylistIsEmpty          = errors.New("playlist is empty")
	ErrorSongСantBeChanged        = errors.New("song is playing and cannot be changed")
)
