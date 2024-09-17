package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/felipeolliveira/_classes_projects/shortlink_with_db/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

func NewHandler(store store.Store) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(store))
	r.Get("/{code}", handleResolveCode(store))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

func handlePost(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJson(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJson(w, Response{Error: "invalid url passed"}, http.StatusBadRequest)
		}

		code, err := store.SaveShortenedURL(r.Context(), body.URL)
		slog.Info(code)

		if err != nil {
			slog.Error("failed do create code", "error", err)
			sendJson(w, Response{Error: "Algo de errado aconteceu"}, http.StatusInternalServerError)
			return
		}

		type codeResponse struct {
			Code string `json:"code"`
		}
		sendJson(w, Response{Data: codeResponse{Code: code}}, http.StatusCreated)
	}
}

func handleResolveCode(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		foundUrl, err := store.GetFullURL(r.Context(), code)

		if err != nil {
			slog.Error("failed to retrieve code", "error", err)

			if errors.Is(err, redis.Nil) {
				http.Error(w, "Short link não encontrado", http.StatusNotFound)
				return
			}

			sendJson(w, Response{Error: "Algo de errado aconteceu"}, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, foundUrl, http.StatusPermanentRedirect)
	}
}
