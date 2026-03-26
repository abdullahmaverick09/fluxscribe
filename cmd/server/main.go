package main

import (
	"log/slog"
	"os"

	"github.com/abdullahmaverick09/fluxscribe/internal/config"
	"github.com/abdullahmaverick09/fluxscribe/internal/database"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()

	if err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	db, err := database.Connect(cfg)

	if err != nil {
		logger.Error("Failed to connect to database ...", "error", err)
		os.Exit(1)
	}

	defer database.Close(db)

	logger.Info("Fluxscribe server is starting ...", "port", cfg.Port, "environment", cfg.Env, "database", cfg.DBName)

	logger.Info("Database connection is established successfully ...")

	logger.Info("Configuration loaded successfully")

}
