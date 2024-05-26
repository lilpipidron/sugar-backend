package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lilpipidron/sugar-backend/internal/config"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/handlers/user"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/middleware/logger"
	"github.com/lilpipidron/sugar-backend/internal/storage/postgresql"
	ur "github.com/lilpipidron/sugar-backend/internal/storage/user"
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	storage, err := postgresql.New(psqlInfo, cfg.DBName)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}

	userRepository := ur.NewUserRepository(storage.DB)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(logger.New(log))
	router.Get("/user", user.NewUserGetter(log, userRepository))
	router.Post("/user", user.NewUserSaver(log, userRepository))
	
	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

}
