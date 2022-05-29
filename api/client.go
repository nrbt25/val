package api

import (
	"fmt"
	"io"
	"net/http"
)

const (
	HOST  string = "https://ap.api.riotgames.com"
	ROUTE string = "val/content/v1/contents"
)

type ClientRequest struct {
	Host  string
	Route string
}

func NewClient(apiKey string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", HOST+"/"+ROUTE, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Riot-Token", apiKey)
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	return body
}
