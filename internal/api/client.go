package api

// Client interacts with 3-rd party joke API
type Client interface {
	// Метод, который возвращает шутки
	// GetJoke returns one joke
	GetJoke() (*JokeResponse, error)
}