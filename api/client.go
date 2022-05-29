package api

import (
	"io"
	"net/http"
)

const (
	HOST string = "https://ap.api.riotgames.com"
)

type ClientRequest struct {
	Host  string
	Route string
}

func NewClient(apiKey, route string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", HOST+"/"+route, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Riot-Token", apiKey)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(string(rune(res.StatusCode)) + ": " + GetStatus(res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return body
}
