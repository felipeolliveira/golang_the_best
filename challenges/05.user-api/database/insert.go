package database

import (
	"log/slog"

	"github.com/felipeolliveira/user_api_golang_challenge/user"
	"github.com/google/uuid"
)

func Insert(firstName string, lastName string, bio string) (id uuid.UUID, err error) {
	mu.Lock()
	defer mu.Unlock()

	id, err = uuid.NewV7()

	if err != nil {
		slog.Error("Database Insert error:", "uuid", err)
		return [16]byte{}, ErrCouldNotSaveUserInDB
	}

	db[id] = user.NewUser(
		firstName,
		lastName,
		bio,
	)

	return id, nil
}
