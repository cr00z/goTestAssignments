package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/repository"
	"strconv"
	"strings"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r PostgresRepository) CreateSong(song *song.Song) (uint64, error) {
	row := r.pool.QueryRow(context.Background(),
		"INSERT INTO playlist (name, duration) VALUES ($1, $2) RETURNING id",
		song.Name, song.Duration,
	)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return 0, errors.Errorf("unable to insert: %v", err)
	}

	return id, nil
}

func (r PostgresRepository) ReadSong(
	ctx context.Context,
	ids []uint64,
	f func(song *song.Song) error,
) error {
	query := "SELECT id, name, duration FROM playlist"
	if len(ids) > 0 {
		idsStr := make([]string, 0)
		for _, id := range ids {
			idsStr = append(idsStr, strconv.FormatUint(id, 10))
		}
		query += " WHERE id IN (" + strings.Join(idsStr, ",") + ")"
	}

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return repository.ErrorEmptyResult
		}
		return err
	}
	defer rows.Close()

	for rows.Next() {
		song := song.Song{}
		rows.Scan(&song.Id, &song.Name, &song.Duration)
		err = f(&song)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r PostgresRepository) UpdateSong(song *song.Song) error {
	//TODO implement me
	return nil
}

func (r PostgresRepository) DeleteSong(id int64) error {
	//TODO implement me
	return nil
}
