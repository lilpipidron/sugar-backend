package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lilpipidron/sugar-backend/internal/config"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/middleware/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *log.Logger {
	var logger *log.Logger

	switch env {
	case envLocal:
		logger = log.NewWithOptions(os.Stdout, log.Options{Level: log.DebugLevel})
	case envDev:
		logger = log.NewWithOptions(os.Stdout, log.Options{Level: log.DebugLevel})
	case envProd:
		logger = log.NewWithOptions(os.Stdout, log.Options{Level: log.DebugLevel})
	}

	return logger
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file", "err", err)
	}

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With("env", cfg.Env)

	log.Info("initializing server", "address", cfg.Address)
	log.Debug("logger debug mode enabled")

	// TODO initialize db

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(logger.New(log))
}
