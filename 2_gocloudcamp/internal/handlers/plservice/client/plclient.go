package client

import (
	"context"
	"gocloudcamp/pkg/plservice"
	"google.golang.org/grpc"
	"log"
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

func (c *PlaylistClient) CreateSong(song *plservice.Song) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &plservice.CreateSongRequest{
		Song: song,
	}

	resp, err := c.service.CreateSong(ctx, req)
	if err != nil {
		log.Fatal("cannot create song request: ", err)
		return
	}

	log.Print("song created with id: ", resp.GetId())
}

func (c *PlaylistClient) ReadSong(ctx context.Context, in *plservice.ReadSongRequest, opts ...grpc.CallOption) (plservice.PlaylistService_ReadSongClient, error) {
	//TODO implement me
	panic("implement me")
}

func (c *PlaylistClient) UpdateSong(ctx context.Context, in *plservice.UpdateSongRequest, opts ...grpc.CallOption) (*plservice.UpdateSongResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *PlaylistClient) DeleteSong(ctx context.Context, in *plservice.DeleteSongRequest, opts ...grpc.CallOption) (*plservice.DeleteSongResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *PlaylistClient) Control(ctx context.Context, in *plservice.ControlRequest, opts ...grpc.CallOption) (*plservice.ControlResponse, error) {
	//TODO implement me
	panic("implement me")
}
