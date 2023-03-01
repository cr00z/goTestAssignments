package server

import (
	"context"
	repository "gocloudcamp/internal/repository/postgres"
	"gocloudcamp/pkg/plservice"
	"log"
)

type PlaylistServer struct {
	repo repository.Repository
	plservice.UnimplementedPlaylistServiceServer
}

func NewPlaylistServer(repo repository.Repository) *PlaylistServer {
	return &PlaylistServer{
		repo: repo,
	}
}

func (s *PlaylistServer) CreateSong(
	ctx context.Context, request *plservice.CreateSongRequest,
) (*plservice.CreateSongResponse, error) {

	if err := contextError(ctx); err != nil {
		return nil, err
	}

	song := request.GetSong()
	log.Print("received create song request with id: ", song.GetId())

	//TODO implement me

	return &plservice.CreateSongResponse{
		Id: 777, // TODO
	}, nil
}

func (s *PlaylistServer) ReadSong(
	request *plservice.ReadSongRequest, server plservice.PlaylistService_ReadSongServer,
) error {

	//TODO implement me

	return nil
}

func (s *PlaylistServer) UpdateSong(
	ctx context.Context, request *plservice.UpdateSongRequest,
) (*plservice.UpdateSongResponse, error) {

	//TODO implement me

	return &plservice.UpdateSongResponse{
		Id: 777, // TODO
	}, nil
}

func (s *PlaylistServer) DeleteSongFromPlaylist(
	ctx context.Context, request *plservice.DeleteSongRequest,
) (*plservice.DeleteSongResponse, error) {

	//TODO implement me

	return &plservice.DeleteSongResponse{
		Id: 777, // TODO
	}, nil
}

func (s *PlaylistServer) Control(
	ctx context.Context, request *plservice.ControlRequest,
) (*plservice.ControlResponse, error) {

	//TODO implement me

	return &plservice.ControlResponse{
		Status: "OK", // TODO
	}, nil
}
