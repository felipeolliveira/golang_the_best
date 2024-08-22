package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/felipeolliveira/shortlink_go/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("All systems offline")
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		return err
	}

	apiKey := os.Getenv("OMDB_KEY")
	handler := api.NewHandler(apiKey)

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
