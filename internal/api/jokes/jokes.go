package jokes

import (
	"encoding/json"
	"example/internal/api"
	"net/http"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil { // Если произошла техническая ошибка
		return nil, err
	} else if resp.StatusCode != http.StatusOK { // Если статус код респонса != 200 (StatusOK)
		return nil, fmt.ErrorF("API request error: %v", err)
	}

	var data api.JokeResponse

	// Ответ приходит json. Мы должны его считать.
	// Т.е. он приходит как набор байт, а в них закодирован json
	err = json.NewDecoder(resp.Body).Decode(&data) // & - это указатель
	if err != nil {
		return nil, err
	}

	return *data, nil
}
