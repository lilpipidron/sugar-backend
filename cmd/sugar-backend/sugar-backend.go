package main

import (
	"log/slog"
	"os"

	"github.com/lilpipidron/sugar-backend/GolandProjects/sugar-backend/internal/config"
	"github.com/lilpipidron/sugar-backend/go/pkg/mod/github.com/labstack/gommon@v0.4.0/log"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

  log := setupLogger(cfg.Env)
log = log.With(slog.String("env", cfg.Env))

  log.Info("initializing server", slog.String("address", cfg.Address))
log.Debug("logger debug mode enabled")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
