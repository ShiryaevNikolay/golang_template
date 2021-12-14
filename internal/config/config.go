package config

// Конфиг сервера
// Server config
type Server struct {
	Port string `yaml:"port" env:"PORT"` // `...` - теги
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0"`

	JokeUrl string `yaml:"joke-url" env:"JOKE_URL"`
}
