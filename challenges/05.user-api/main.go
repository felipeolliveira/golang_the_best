package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/felipeolliveira/user_api_golang_challenge/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Error when running app:", "err", err)
		os.Exit(1)
	}

	slog.Info("All system has been running")
}

func run() error {
	handler := api.NewHandler()
	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
