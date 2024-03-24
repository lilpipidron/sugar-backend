package config

import (
	"log"
	"os"
)

type Config struct {
	Env        string `yaml:"env" env-default:"development"`
	HTTPServer `yaml:"http_server"`
	Database   `yaml:"database"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0:8080"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	DBName   string `yaml:"dbname" env-default:"postgres"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
  if err != nil {
    log.Fatalf("error reading config file: %s", err)
  }

  return &cfg
}
