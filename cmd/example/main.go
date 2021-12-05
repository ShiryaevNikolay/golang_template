package main

import (
	"example/internal/handler"

	"github.com/go-chi/chi"
)

func main() {

	h :=  handler.NewHandler() // handler, который будет слушать запросы

	r := chi.NewRouter() // Создается роутер. Он обрабатывает запросы (GET, POST, DELETE...)

	// Запрос типа GET.
	//
	// Первым параметром принимает в себя путь. 
	// Второй параметр - функция handler, 
	// который будет слушать этот запрос. В данном случае это функция Hello, 
	// которая находится в файле handler.go
	r.Get("/hello", h.Hello)
}
