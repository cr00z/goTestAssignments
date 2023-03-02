package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"gocloudcamp/internal/domain/playlist"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/handlers/plservice/server"
	"gocloudcamp/internal/player"
	repository "gocloudcamp/internal/repository/postgres"
	"gocloudcamp/pkg/helpers/postgres"
	"gocloudcamp/pkg/plservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	serverAddr = "0.0.0.0:50051"
)

var pgConfig = &postgres.Config{
	Host:     "127.0.0.1",
	Port:     "5432",
	Username: "user",
	Password: "password",
	DBName:   "playlist",
	Timeout:  5,
	MaxConns: 5,
}

func main() {
	// db

	poolConfig, err := postgres.NewPoolConfig(pgConfig)
	if err != nil {
		log.Fatal("pool config error: ", err)
	}

	pool, err := postgres.NewPool(poolConfig)
	if err != nil {
		log.Fatal("connect to db failed: ", err)
	}

	log.Print("db connected")

	repo := repository.NewPostgresRepository(pool)

	// goose

	gooseDB, err := sql.Open("postgres", poolConfig.ConnString())
	if err != nil {
		log.Fatal("goose connect to db failed: ", err)
	}

	err = gooseDB.Ping()
	if err != nil {
		log.Fatal("goose connect to db failed: ", err)
	}

	err = goose.Up(gooseDB, "./migrations")

	// playlist

	pl := playlist.NewPlaylist()
	err = repo.ReadSong(context.Background(), []uint64{}, func(song *song.Song) error {
		pl.AddSong(song)
		return nil
	})
	if err != nil {
		log.Fatal("playlist start failed: ", err)
	}
	go player.Start(pl)

	log.Print("playlist started")

	// grpc

	grpcServer := grpc.NewServer()
	playlistService := server.NewPlaylistServer(repo, pl)
	plservice.RegisterPlaylistServiceServer(grpcServer, playlistService)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
