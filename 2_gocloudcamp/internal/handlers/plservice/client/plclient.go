package client

import (
	"context"
	"github.com/pkg/errors"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/pkg/plservice"
	"google.golang.org/grpc"
	"io"
	"time"
)

const requestTimeout = time.Second * 5

type PlaylistClient struct {
	service plservice.PlaylistServiceClient
}

func NewPlaylistClient(cc *grpc.ClientConn) *PlaylistClient {
	service := plservice.NewPlaylistServiceClient(cc)
	return &PlaylistClient{
		service: service,
	}
}

func (c *PlaylistClient) CreateSong(song *song.Song) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &plservice.CreateSongRequest{
		Song: &plservice.Song{
			Name:     song.Name,
			Duration: uint64(song.Duration),
		},
	}

	resp, err := c.service.CreateSong(ctx, req)
	if err != nil {
		return 0, errors.Errorf("cannot execute create song request: %v", err)
	}

	return resp.GetId(), nil
}

func (c *PlaylistClient) ReadSong(ids []uint64) ([]*song.Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &plservice.ReadSongRequest{
		Id: ids,
	}

	stream, err := c.service.ReadSong(ctx, req)
	if err != nil {
		return nil, errors.Errorf("cannot execute read song request: %v", err)
	}

	songs := make([]*song.Song, 0)
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return songs, nil
			}
			return nil, errors.Errorf("cannot receive response: %v", err)
		}

		resSong := res.GetSong()
		songs = append(songs, &song.Song{
			Id:       resSong.GetId(),
			Name:     resSong.GetName(),
			Duration: time.Duration(resSong.GetDuration()),
		})
	}
	//return songs, nil
}

func (c *PlaylistClient) UpdateSong(song *song.Song) (uint64, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &plservice.UpdateSongRequest{
		Song: &plservice.Song{
			Id:       song.Id,
			Name:     song.Name,
			Duration: uint64(song.Duration),
		},
	}

	resp, err := c.service.UpdateSong(ctx, req)
	if err != nil {
		return 0, errors.Errorf("cannot execute update song request: %v", err)
	}

	return resp.GetId(), nil
}

func (c *PlaylistClient) DeleteSong(id uint64) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &plservice.DeleteSongRequest{
		Id: id,
	}

	resp, err := c.service.DeleteSong(ctx, req)
	if err != nil {
		return 0, errors.Errorf("cannot execute delete song request: %v", err)
	}

	return resp.GetId(), nil
}

func (c *PlaylistClient) Control(ctrl plservice.ControlRequest_Action) (string, error) {
	//TODO implement me
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &plservice.ControlRequest{
		Action: ctrl,
	}

	resp, err := c.service.Control(ctx, req)
	if err != nil {
		return "", errors.Errorf("cannot execute control request: %v", err)
	}

	return resp.GetStatus(), nil
}
