package api

import (
	"encoding/json"
	"val/api"
)

type MatchDto struct {
	MatchInfo    MatchInfoDto     `json:"matchInfo"`
	Players      []PlayerDto      `json:"players"`
	Coaches      []CoachDto       `json:"coaches"`
	Teams        []TeamDto        `json:"teams"`
	RoundResults []RoundResultDto `json:"roundResults"`
}

type MatchInfoDto struct {
	MatchId            string `json:"matchId"`
	MapId              string `json:"mapId"`
	GameLengthMillis   int    `json:"gameLengthMillis"`
	GameStartMillis    int64  `json:"gameStartMillis"`
	ProvisioningFlowId string `json:"provisioningFlowId"`
	IsCompleted        bool   `json:"isCompleted"`
	CustomGameName     string `json:"customGameName"`
	QueueId            string `json:"queueId"`
	GameMode           string `json:"gameMode"`
	IsRanked           bool   `json:"isRanked"`
	SeasonId           string `json:"seasonId"`
}

type PlayerDto struct {
	Puuid           string         `json:"puuid"`
	GameName        string         `json:"gameName"`
	TagLine         string         `json:"tagLine"`
	TeamId          string         `json:"teamId"`
	PartyId         string         `json:"partyId"`
	CharacterId     string         `json:"characterId"`
	Stats           PlayerStatsDto `json:"stats"`
	CompetitiveTier int            `json:"competitiveTier"`
	PlayerCard      string         `json:"playerCard"`
	PlayerTitle     string         `json:"playerTitle"`
}

type PlayerStatsDto struct {
	Score          int             `json:"score"`
	RoundPlayed    int             `json:"roundPlayed"`
	Kills          int             `json:"kills"`
	Deaths         int             `json:"deaths"`
	Assists        int             `json:"assists"`
	PlayTimeMillis int             `json:"playTimeMillis"`
	AbilityCasts   AbilityCastsDto `json:"abilityCasts"`
}

type AbilityCastsDto struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

type CoachDto struct {
	Puuid  string `json:"puuid"`
	TeamId string `json:"teamId"`
}

type TeamDto struct {
	TeamId       string `json:"teamId"`
	Win          bool   `json:"win"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

type RoundResultDto struct {
	RoundNum              int                   `json:"roundNum"`
	RoundResult           string                `json:"roundResult"`
	RoundCeremony         string                `json:"roundCeremony"`
	WinningTeam           string                `json:"winningTeam"`
	BombPlanter           string                `json:"bombPlanter"`
	BombDefuser           string                `json:"bombDefuser"`
	PlantRoundTime        int                   `json:"plantRoundTime"`
	PlantPlayerLocations  []PlayerLocationsDto  `json:"plantPlayerLocations"`
	PlantLocation         LocationDto           `json:"plantLocation"`
	PlantSite             string                `json:"plantSite"`
	DefuseRoundTime       int                   `json:"defuseRoundTime"`
	DefusePlayerLocations []PlayerLocationsDto  `json:"defusePlayerLocation"`
	PlayerStats           []PlayerRoundStatsDto `json:"playerStats"`
	RoundResultCode       string                `json:"roundResultCode"`
}

type PlayerLocationsDto struct {
	Puuid       string      `json:"puuid"`
	ViewRadians int64       `json:"viewRadians"`
	Location    LocationDto `json:"location"`
}

type LocationDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerRoundStatsDto struct {
	Puuid   string      `json:"puuid"`
	Kills   []KillDto   `json:"kills"`
	Damage  []DamageDto `json:"damageDto"`
	Score   int         `json:"score"`
	Economy EconomyDto  `json:"economyDto"`
	Ability AbilityDto  `json:"ability"`
}

type KillDto struct {
	TimeSinceGameStartMillis  int64                `json:"timeSinceGameStartMillis"`
	TimeSinceRoundStartMillis int64                `json:"timeSinceRoundStartMillis"`
	Killer                    string               `json:"killer"`
	Victim                    string               `json:"victim"`
	VictimLocation            LocationDto          `json:"victimLocation"`
	Assistants                []string             `json:"assistants"`
	PlayerLocations           []PlayerLocationsDto `json:"playerLocations"`
	FinishingDamage           FinishingDamageDto   `json:"finishingDamageDto"`
}

type FinishingDamageDto struct {
	DamageType          string `json:"damageType"`
	DamageItem          string `json:"damageItem"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
}

type DamageDto struct {
	Receiver  string `json:"receiver"`
	Damage    int    `json:"damage"`
	LegShots  int    `json:"legShots"`
	BodyShots int    `json:"bodyShots"`
	HeadShots int    `json:"headShots"`
}

type EconomyDto struct {
	LoadoutValue int    `json:"loadoutValue"`
	Weapon       string `json:"weapon"`
	Armor        string `json:"armor"`
	Remaining    int    `json:"remaining"`
	Spent        int    `json:"spent"`
}

type AbilityDto struct {
	GrenadeEffects  string `json:"grenadeEffects"`
	Ability1Effects string `json:"ability1Effects"`
	Ability2Effects string `json:"ability2Effects"`
	UltimateEffects string `json:"ultimateEffects"`
}

func NewValMatch(apiKey string, matchId int) MatchDto {
	client := api.NewClient(apiKey, "val/match/v1/matches/"+string(rune(matchId)))

	var valMatch = MatchDto{}

	json.Unmarshal(client, &valMatch)

	return valMatch
}
