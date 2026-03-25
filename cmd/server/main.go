package main

import (
	"log/slog"
	"os"

	"github.com/abdullahmaverick09/fluxscribe/internal/config"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	cfg, err := config.Load()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	logger.Info("Fluxscribe server is starting ...", "port", cfg.Port, "environment", cfg.Env, "database", cfg.DBName)

	logger.Info("Configuration loaded successfully")

}
