package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"storage_path" env-default:"./storage/storage.db" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"address" env-default:"localhost:8082" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env:"timeout" env-default:"4s" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"idle_timeout" env-default:"60s" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./confing/local.yaml"
	}

	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %s", err)

	}

	return &cfg
}
