package repository

import (
	"context"
	"errors"
	"gocloudcamp/internal/domain/song"
)

var ErrorEmptyResult = errors.New("empty result")

type Repository interface {
	CreateSong(song *song.Song) (uint64, error)
	ReadSong(ctx context.Context, ids []uint64, f func(song *song.Song) error) error
	UpdateSong(song *song.Song) (uint64, error)
	DeleteSong(id uint64) error
}
