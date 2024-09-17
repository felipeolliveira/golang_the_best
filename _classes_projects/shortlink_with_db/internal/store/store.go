package store

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type store struct {
	db *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

func NewStore(db *redis.Client) Store {
	return store{db}
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string

	for range 5 {
		code = genCode()

		if err := s.db.HGet(ctx, "encurtador", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
		}
	}

	slog.Info(code)

	if err := s.db.HSet(ctx, "encurtador", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code in encurtador hashmap: %w", err)
	}

	return code, nil
}

func (s store) GetFullURL(ctx context.Context, code string) (string, error) {
	fullUrl, err := s.db.HGet(ctx, "encurtador", code).Result()

	if err != nil {
		return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
	}

	return fullUrl, nil
}
