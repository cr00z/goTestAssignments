package server

import (
	"context"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/repository"
	"gocloudcamp/pkg/plservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
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

	reqSong := request.GetSong()
	log.Print("received create song request with id: ", reqSong.GetId())

	id, err := s.repo.CreateSong(&song.Song{
		Name:     reqSong.Name,
		Duration: time.Duration(reqSong.Duration),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot save song to the repo: %v", err)
	}

	return &plservice.CreateSongResponse{
		Id: id,
	}, nil
}

func (s *PlaylistServer) ReadSong(
	request *plservice.ReadSongRequest,
	stream plservice.PlaylistService_ReadSongServer,
) error {
	ids := request.GetId()
	log.Print("received read song request with ids: ", ids)

	err := s.repo.ReadSong(
		stream.Context(),
		ids,
		func(song *song.Song) error {
			res := &plservice.ReadSongResponse{
				Song: &plservice.Song{
					Id:       song.Id,
					Name:     song.Name,
					Duration: uint64(song.Duration),
				},
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("seng song with id: %d", song.Id)
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot read song from the repo: %v", err)
	}

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
