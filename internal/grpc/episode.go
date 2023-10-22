package grpc

import (
	"context"
	proto "github.com/ElladanTasartir/buffy-grpc/gen/go/episodes/v1"
	buffyClient "github.com/ElladanTasartir/buffy-grpc/internal/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EpisodesService struct {
	client *buffyClient.BuffyClient
}

func NewEpisodesService() (*EpisodesService, error) {
	client, err := buffyClient.NewBuffyClient()
	if err != nil {
		return nil, err
	}

	return &EpisodesService{
		client: client,
	}, nil
}

func (es *EpisodesService) GetEpisode(ctx context.Context, req *proto.EpisodeRequest) (*proto.EpisodeResponse, error) {
	episode, err := es.client.GetEpisode(int32(req.Series), req.Season, req.Episode)
	if err != nil {
		err = status.New(codes.NotFound, err.Error()).Err()
		return nil, err
	}

	return &proto.EpisodeResponse{
		Episode: &proto.Episode{
			Id:          episode.Id,
			Name:        episode.Episode,
			Description: episode.Description,
			Screenshot:  episode.Screenshot,
			Trivia:      episode.Trivia,
		},
	}, nil
}

func (es *EpisodesService) GetSeason(ctx context.Context, req *proto.SeasonRequest) (*proto.SeasonResponse, error) {
	episodes, err := es.client.GetSeason(int32(req.Series), req.Season)
	if err != nil {
		err = status.New(codes.NotFound, err.Error()).Err()
		return nil, err
	}

	var response []*proto.Episode
	for _, episode := range episodes {
		response = append(response, &proto.Episode{
			Id:          episode.Id,
			Name:        episode.Episode,
			Description: episode.Description,
			Screenshot:  episode.Screenshot,
			Trivia:      episode.Trivia,
		})
	}

	return &proto.SeasonResponse{
		Episodes: response,
	}, nil
}
