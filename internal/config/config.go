package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string  `yaml:"env" env:"ENV" env-default:"development"`
	Storage    Storage `yaml:"storage" env-required:"true"`
	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"0.0.0.0:8080"`
		Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
		Password    string        `yaml:"password" env-required:"true" env-default:"hla"`
		User        string        `yaml:"user" env-required:"true" env-default:"hla"`
	} `yaml:"http_server"`
}

type Storage struct {
	Path string `yaml:"path" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH переменная окружения не установлена")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("Ошибка открытия конфигурационного файла: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла: %s", err)
	}
	return &cfg
}
