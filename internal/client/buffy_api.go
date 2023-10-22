package client

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

const API_URL = "https://buffy-angel-api.onrender.com/api/buffy"

type BuffyClient struct {
	url        string
	httpClient http.Client
	logger     *zap.Logger
}

type BuffyEpisode struct {
	Id          string `json:"_id"`
	Episode     string `json:"episodeName"`
	Description string `json:"description"`
	Trivia      string `json:"trivia"`
	Screenshot  string `json:"episodeScreenshot"`
}

func NewBuffyClient() (*BuffyClient, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return &BuffyClient{}, err
	}

	return &BuffyClient{
		url:        API_URL,
		httpClient: http.Client{},
		logger:     logger,
	}, nil
}

func (client *BuffyClient) GetEpisode(season, episode int32) (BuffyEpisode, error) {
	resp, err := client.httpClient.Get(fmt.Sprintf("%s/season/%d/%d", client.url, season, episode))
	if err != nil {
		return BuffyEpisode{}, err
	}

	defer resp.Body.Close()

	var buffyEpisode []BuffyEpisode
	err = json.NewDecoder(resp.Body).Decode(&buffyEpisode)
	if err != nil {
		return BuffyEpisode{}, err
	}

	if len(buffyEpisode) < 1 {
		return BuffyEpisode{}, fmt.Errorf("there's no episode from season %d episode %d", season, episode)
	}

	return buffyEpisode[0], err
}

func (client *BuffyClient) GetSeason(season int32) ([]BuffyEpisode, error) {
	var episodes []BuffyEpisode
	resp, err := client.httpClient.Get(fmt.Sprintf("%s/season/%d", client.url, season))
	if err != nil {
		return episodes, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&episodes)
	if err != nil {
		return episodes, err
	}

	return episodes, nil
}
