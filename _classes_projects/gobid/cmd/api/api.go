package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/api"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	gob.Register(uuid.UUID{})
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		os.Getenv("GOBID_DATABASE_HOST"),
		os.Getenv("GOBID_DATABASE_PORT"),
		os.Getenv("GOBID_DATABASE_USER"),
		os.Getenv("GOBID_DATABASE_PASSWORD"),
		os.Getenv("GOBID_DATABASE_NAME"),
	)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(pool)
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	ws := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if os.Getenv("GOBID_CORS") == "true" {
				// This variable should be fill with origins to allow
				allowOrigins := []string{}
				origin := r.Header.Get("Origin")

				return slices.Contains(allowOrigins, origin)
			}

			//
			return true
		},
	}

	api := api.Api{
		Router:         chi.NewMux(),
		UserService:    services.NewUserService(pool),
		ProductService: services.NewProductService(pool),
		BidsService:    services.NewBidsService(pool),
		Session:        sessionManager,
		WsUpgrader:     ws,
		AuctionLobby:   services.NewAuctionLobby(),
	}

	api.BindRoutes()

	port := os.Getenv("GOBID_APP_PORT")
	if port == "" {
		port = "8080"
	}
	serverPort := fmt.Sprintf(":%s", port)

	fmt.Printf("Starting server on port %s\n", serverPort)
	if err := http.ListenAndServe(serverPort, api.Router); err != nil {
		panic(err)
	}
}
