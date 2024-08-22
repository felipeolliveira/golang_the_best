package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", handleGetUsers)
		r.Post("/users", handleCreateUser)

		r.Get("/users/{id}", handleGetUserById)
		r.Put("/users/{id}", handleUpdateUserById)
		r.Delete("/users/{id}", handleRemoveUserById)
	})

	return r
}
