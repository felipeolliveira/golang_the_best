package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/felipeolliveira/_classes_projects/shortlink_with_db/internal/api"
	"github.com/felipeolliveira/_classes_projects/shortlink_with_db/internal/store"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}

	slog.Info("All systems offline")
}

func run() error {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	store := store.NewStore(db)
	handler := api.NewHandler(store)

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
