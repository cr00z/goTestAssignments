package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leporo/sqlf"
	"github.com/pkg/errors"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/repository"
	"strconv"
	"strings"
)

const TableName = "playlist"

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r PostgresRepository) CreateSong(song *song.Song) (uint64, error) {
	const query = "INSERT INTO " + TableName + " (name, duration) VALUES ($1, $2) RETURNING id"
	row := r.pool.QueryRow(context.Background(), query, song.Name, song.Duration)

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
	var query = "SELECT id, name, duration FROM " + TableName
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
		err = rows.Scan(&song.Id, &song.Name, &song.Duration)
		if err != nil {
			return err
		}
		err = f(&song)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r PostgresRepository) UpdateSong(song *song.Song) (uint64, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.Update(TableName)
	if song.Name != "" {
		b = b.Set("name", song.Name)
	}
	if song.Duration != 0 {
		b = b.Set("duration", song.Duration)
	}
	b.Where("id=?", song.Id).Returning("id")

	row := r.pool.QueryRow(context.Background(), b.String(), b.Args()...)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return 0, errors.Errorf("unable to update: %v", err)
	}

	return id, nil
}

func (r PostgresRepository) DeleteSong(id uint64) error {
	const query = "DELETE FROM " + TableName + " WHERE id = $1"

	ct, err := r.pool.Exec(context.Background(), query, id)
	if err != nil {
		return errors.Errorf("unable to delete: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return repository.ErrorEmptyResult
	}

	return nil
}
