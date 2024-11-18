package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	queries *pgstore.Queries
	pool    *pgxpool.Pool
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var UserServiceErr = struct {
	ErrEmailOrUserNameAlreadyExists error
	ErrInvalidCredentials           error
	ErrCouldNotCreateUser           error
	ErrCouldNotAuthenticate         error
}{
	errors.New("email or username already exists"),
	errors.New("invalid credentials"),
	errors.New("could not create user"),
	errors.New("could not authenticate user"),
}

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		slog.Error("user service", "create user", err)
		return uuid.UUID{}, UserServiceErr.ErrCouldNotCreateUser
	}

	args := pgstore.CreateUserParams{
		UserName:     userName,
		Email:        email,
		PasswordHash: hash,
		Bio:          bio,
	}

	id, err := us.queries.CreateUser(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, UserServiceErr.ErrEmailOrUserNameAlreadyExists
		}

		slog.Error("user service", "create user", err)
		return uuid.UUID{}, UserServiceErr.ErrCouldNotCreateUser
	}

	return id, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	user, err := us.queries.GetUserByEmail(ctx, email)
	if err != nil {
		slog.Error("user service", "authenticate user", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, UserServiceErr.ErrInvalidCredentials
		}

		slog.Error("user service", "authenticate user", err)
		return uuid.UUID{}, UserServiceErr.ErrCouldNotAuthenticate
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, UserServiceErr.ErrInvalidCredentials
		}

		slog.Error("user service", "authenticate user", err)
		return uuid.UUID{}, UserServiceErr.ErrCouldNotAuthenticate
	}

	return user.ID, nil
}
