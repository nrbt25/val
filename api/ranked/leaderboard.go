package ranked

import (
	"encoding/json"
	"val/client"
)

type LeaderBoardDto struct {
	Shard        string      `json:"shard"`
	ActId        string      `json:"actId"`
	TotalPLayers int64       `json:"totalPlayers"`
	Players      []PlayerDto `json:"players"`
}

type PlayerDto struct {
	Puuid           string `json:"puuid"`
	GameName        string `json:"gameName"`
	TagLine         string `json:"tagLine"`
	LeaderboardRank int64  `json:"leaderboardRank"`
	RankedRating    int64  `json:"rankedRating"`
	NumberOfWins    int64  `json:"numberOfWins"`
}

func GetLeaderBoard(apiKey, actId string) LeaderBoardDto {
	client := client.NewClient(apiKey, "val/match/v1/matches/"+actId)

	var leaderBoard = LeaderBoardDto{}

	json.Unmarshal(client, &leaderBoard)

	return leaderBoard
}
