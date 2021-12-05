package main

// Принято группировать импорты
// 
// 1. Стандартные импорты
// 2. Внешние импорты
// 3. Пути внутри проекта
import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"example/internal/handler"
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


	err := http.ListenAndServe(":8080", r) // Возвращает ошибку
	log.Fatal(err)
}
