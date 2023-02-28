package main

import (
	repository "gocloudcamp/internal/repository/postgres"
	"gocloudcamp/pkg/helpers/postgres"
	"gocloudcamp/pkg/plserver"
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

	// grpc

	grpcServer := grpc.NewServer()
	playlistService := plserver.NewPlaylistServer(repo)
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
