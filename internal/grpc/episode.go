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

	directors, writers, cast := es.mapCastMembersToResponse(&episode)

	return &proto.EpisodeResponse{
		Episode: &proto.Episode{
			Id:          episode.Id,
			Name:        episode.Episode,
			Description: episode.Description,
			Screenshot:  episode.Screenshot,
			Trivia:      episode.Trivia,
			Directors:   directors,
			Writers:     writers,
			Cast:        cast,
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
		directors, writers, cast := es.mapCastMembersToResponse(&episode)
		response = append(response, &proto.Episode{
			Id:          episode.Id,
			Name:        episode.Episode,
			Description: episode.Description,
			Screenshot:  episode.Screenshot,
			Trivia:      episode.Trivia,
			Directors:   directors,
			Writers:     writers,
			Cast:        cast,
		})
	}

	return &proto.SeasonResponse{
		Episodes: response,
	}, nil
}

func (es *EpisodesService) mapCastMembersToResponse(episode *buffyClient.BuffyEpisode) ([]*proto.CastMember, []*proto.CastMember, []*proto.CastMember) {
	var cast []*proto.CastMember
	for _, member := range episode.Cast {
		cast = append(cast, &proto.CastMember{
			Id:      member.Id,
			Name:    member.Name,
			Picture: member.Picture,
		})
	}
	var directors []*proto.CastMember
	for _, director := range episode.Directors {
		directors = append(directors, &proto.CastMember{
			Id:      director.Id,
			Name:    director.Name,
			Picture: director.Picture,
		})
	}
	var writers []*proto.CastMember
	for _, writer := range episode.Writers {
		writers = append(writers, &proto.CastMember{
			Id:      writer.Id,
			Name:    writer.Name,
			Picture: writer.Picture,
		})
	}

	return directors, writers, cast
}
