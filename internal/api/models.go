package api

// JokeResponse is joke API response type
// Описание типа, который присылает
type JokeResponse struct {
	Joke string `json:"joke"` // можно было явно не прописывать тег. "joke" это ключ
}