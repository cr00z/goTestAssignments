package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	Timeout  int
	MaxConns int32
}

func NewPoolConfig(cfg *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Timeout,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = cfg.MaxConns

	return poolConfig, nil
}

func NewPool(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	_, err = pool.Exec(context.Background(), ";")
	if err != nil {
		return nil, err
	}

	return pool, nil
}
