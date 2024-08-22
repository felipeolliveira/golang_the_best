package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/felipeolliveira/shortlink_go/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("All systems offline")
}

func run() error {
	db := make(map[string]string)
	handler := api.NewHandler(db)

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
