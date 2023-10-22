package grpc

import (
	"context"
	proto "github.com/ElladanTasartir/buffy-grpc/gen/go/episodes/v1"
	buffyClient "github.com/ElladanTasartir/buffy-grpc/internal/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EpisodesService struct{}

func (es EpisodesService) GetEpisode(ctx context.Context, req *proto.EpisodeRequest) (*proto.EpisodeResponse, error) {
	client, err := buffyClient.NewBuffyClient()
	if err != nil {
		err = status.New(codes.Unknown, err.Error()).Err()
		return nil, err
	}

	episode, err := client.GetEpisode(req.Season, req.Episode)
	if err != nil {
		err = status.New(codes.NotFound, err.Error()).Err()
		return nil, err
	}

	return &proto.EpisodeResponse{
		Id:          episode.Id,
		Name:        episode.Episode,
		Description: episode.Description,
	}, nil
}
