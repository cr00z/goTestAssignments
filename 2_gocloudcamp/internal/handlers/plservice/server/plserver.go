package server

import (
	"context"
	"gocloudcamp/internal/domain/playlist"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/repository"
	"gocloudcamp/pkg/plservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type PlaylistServer struct {
	repo     repository.Repository
	playlist *playlist.Playlist
	plservice.UnimplementedPlaylistServiceServer
}

func NewPlaylistServer(repo repository.Repository, pl *playlist.Playlist) *PlaylistServer {
	return &PlaylistServer{
		repo:     repo,
		playlist: pl,
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

	song := &song.Song{
		Name:     reqSong.Name,
		Duration: time.Duration(reqSong.Duration),
	}
	id, err := s.repo.CreateSong(song)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot save song to the repo: %v", err)
	}

	song.Id = id
	s.playlist.AddSong(song)

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
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	id := request.GetSong().GetId()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "unknown id")
	}

	name := request.GetSong().GetName()
	duration := request.GetSong().GetDuration()
	if name == "" && duration == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "data not changed")
	}

	log.Print("received update song request with id: ", id)

	updSong := &song.Song{
		Id:       id,
		Name:     name,
		Duration: time.Duration(duration),
	}

	err := s.playlist.ModifySong(updSong)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update song from the playlist: %v", err)
	}

	updatedId, err := s.repo.UpdateSong(updSong)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update song from the repo: %v", err)
	}

	return &plservice.UpdateSongResponse{
		Id: updatedId,
	}, nil
}

func (s *PlaylistServer) DeleteSong(
	ctx context.Context, request *plservice.DeleteSongRequest,
) (*plservice.DeleteSongResponse, error) {
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	id := request.GetId()
	log.Print("received delete song request with id: ", id)

	err := s.playlist.RemoveSong(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot delete song from the playlist: %v", err)
	}

	err = s.repo.DeleteSong(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot delete song from the repo: %v", err)
	}

	return &plservice.DeleteSongResponse{
		Id: id,
	}, nil
}

func (s *PlaylistServer) Control(
	ctx context.Context, request *plservice.ControlRequest,
) (*plservice.ControlResponse, error) {
	result := "OK"
	var err error
	var song *song.Song

	ctrl := request.GetAction()
	switch ctrl {
	case plservice.ControlRequest_PLAY:
		err = s.playlist.Play()
	case plservice.ControlRequest_PAUSE:
		err = s.playlist.Pause()
	case plservice.ControlRequest_NEXT:
		err = s.playlist.Next()
		song, _ = s.playlist.CurrentSong()
		result = "song: " + song.Name
	case plservice.ControlRequest_PREV:
		err = s.playlist.Prev()
		song, _ = s.playlist.CurrentSong()
		result = "song: " + song.Name
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "control error: %v", err)
	}

	return &plservice.ControlResponse{
		Status: result,
	}, nil
}
