package api

import (
	"encoding/json"
	"val/api"
)

type RecentMatchesDto struct {
	CurrentTime int64    `json:"currentTime"`
	MatchIds    []string `json:"matchIds"`
}

func GetRecentMatches(apiKey, queue string) RecentMatchesDto {
	client := api.NewClient(apiKey, "val/match/v1/recent-matches/by-queue/"+queue)

	var recentMatches = RecentMatchesDto{}

	json.Unmarshal(client, &recentMatches)

	return recentMatches
}
