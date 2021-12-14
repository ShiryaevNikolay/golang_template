package handler

import (
	"example/internal/api"
	"fmt"
	"net/http"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.GetJoke() // возвращает joke и ошибку
	if err != nil { // Если ошибка, то нужно ее записать в респонс
		http.Error(w, err.Error(), http.StatusInternalServerError) // принимает ResponseWriter, текст ошибки и статус код
		return
	}

	fmt.Fprint(w, joke.Joke)
}
