package config

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"development"`
	HTTPServer `yaml:"http_server"`
	Database
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0:8080"`
}

type Database struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

func retrieveDatabaseConfig(config *Config) {
	var dbData Database

	dbData.User = os.Getenv("DB_USER")
	dbData.Password = os.Getenv("DB_PASSWORD")
	dbData.DBName = os.Getenv("DB_NAME")
	dbData.Host = os.Getenv("DB_HOST")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("failed to parse database port", "err", err)
	}
	dbData.Port = port

	config.Database = dbData
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("environment variable CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatal("failed to open config file", "err", err)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatal("failed to read config file", "err", err)
	}

	retrieveDatabaseConfig(&config)

	return &config
}
