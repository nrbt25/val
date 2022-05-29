package api

import (
	"encoding/json"
	"val/api"
)

type ContentItemDto struct {
	Name           string            `json:"name"`
	LocalizedNames map[string]string `json:"localizedNames"`
	Id             string            `json:"id"`
	AssetName      string            `json:"assetName"`
}

type LocalizedNamesDto struct {
	ArAE string `json:"ar-AE"`
	DeDE string `json:"de-DE"`
	EnDB string `json:"en-DB"`
	EnUS string `json:"en-US"`
	EsES string `json:"es-ES"`
	EsMX string `json:"es-MX"`
	FrFR string `json:"fr-FR"`
	IdID string `json:"id-ID"`
	ItIT string `json:"it-IT"`
	KaJP string `json:"ja-JP"`
	KoKR string `json:"ko-KR"`
	PlPL string `json:"pl-PL"`
	PtBR string `json:"pt-BR"`
	RuRU string `json:"ru-RU"`
	ThTH string `json:"th-TH"`
	TrTR string `json:"tr-TR"`
	ViVN string `json:"vi-VN"`
	ZhCN string `json:"zh-CN"`
	ZhTW string `json:"zh-TW"`
}

type ActDto struct {
	Name           string            `json:"name"`
	LocalizedNames LocalizedNamesDto `json:"localizedNames"`
	Id             string            `json:"id"`
	IsActive       bool              `json:"bool"`
}

type ValContent struct {
	Version      string           `json:"version"`
	Characters   []ContentItemDto `json:"characters"`
	Maps         []ContentItemDto `json:"maps"`
	Chromas      []ContentItemDto `json:"chromas"`
	Skins        []ContentItemDto `json:"skins"`
	SkinLevels   []ContentItemDto `json:"skinLevels"`
	Equips       []ContentItemDto `json:"equips"`
	GameModes    []ContentItemDto `json:"gameModes"`
	Sprays       []ContentItemDto `json:"sprays"`
	SprayLevels  []ContentItemDto `json:"sprayLevels"`
	Charms       []ContentItemDto `json:"charms"`
	CharmsLevels []ContentItemDto `json:"charmsLevels"`
	PlayerCards  []ContentItemDto `json:"playerCards"`
	PlayerTitles []ContentItemDto `json:"playerTitles"`
	Acts         []ActDto         `json:"acts"`
}

func GetContent(apiKey string) ValContent {
	client := api.NewClient(apiKey, "val/content/v1/contents")

	var valContent = ValContent{}

	json.Unmarshal(client, &valContent)

	return valContent
}
