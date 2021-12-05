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
	"github.com/ilyakaznacheev/cleanenv"

	"example/internal/config"
	"example/internal/handler"
)

func main() {
	cfg := config.Server{} // создание структуры Server из config.go
	
	// Считываем файл config.yml
	// И по тегам, которые описываются в config.go в структуре Server
	// считываем конфиг
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h :=  handler.NewHandler() // handler, который будет слушать запросы

	r := chi.NewRouter() // Создается роутер. Он обрабатывает запросы (GET, POST, DELETE...)

	// Запрос типа GET.
	//
	// Первым параметром принимает в себя путь. 
	// Второй параметр - функция handler, 
	// который будет слушать этот запрос. В данном случае это функция Hello, 
	// которая находится в файле handler.go
	r.Get("/hello", h.Hello)

	path := cfg.Host+":"+cfg.Port

	log.Printf("starting server at %s", path)
	err = http.ListenAndServe(path, r) // Возвращает ошибку
	log.Fatal(err)
	log.Print("shutting server down")
}
