package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"gocloudcamp/internal/domain/song"
)

type Repository interface {
	AddSong(song song.Song) error
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r PostgresRepository) AddSong(song song.Song) error {
	//TODO implement me
	return nil
}
