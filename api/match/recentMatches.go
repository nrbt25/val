package match

import (
	"encoding/json"
	"val/client"
)

type RecentMatchesDto struct {
	CurrentTime int64    `json:"currentTime"`
	MatchIds    []string `json:"matchIds"`
}

func GetRecentMatches(apiKey, queue string) RecentMatchesDto {
	client := client.NewClient(apiKey, "val/match/v1/recent-matches/by-queue/"+queue)

	var recentMatches = RecentMatchesDto{}

	json.Unmarshal(client, &recentMatches)

	return recentMatches
}
