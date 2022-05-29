package api

import (
	"encoding/json"
	"val/api"
)

type MatchListDto struct {
	Puuid   string              `json:"puuid"`
	History []MatchlistEntryDto `json:"mathlistEntryDto"`
}

type MatchlistEntryDto struct {
	MatchId            string `json:"matchId"`
	GameStartTimeMills int64  `json:"gameStartTimeMillis"`
	TeamId             string `json:"teamId"`
}

func GetMatchlistByPuuid(apiKey, puuid string) MatchListDto {
	client := api.NewClient(apiKey, "val/match/v1/matchlists/by-puuid/"+puuid)

	var matchListDto = MatchListDto{}

	json.Unmarshal(client, &matchListDto)

	return matchListDto
}
