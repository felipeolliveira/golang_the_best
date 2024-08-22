package api

import (
	"net/http"

	"github.com/felipeolliveira/shortlink_go/omdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleSearchMovie(apiKey))

	return r
}

func handleSearchMovie(apikey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apikey, search)

		if err != nil {
			sendJson(w, Response{Error: "something wrong with omdb"}, http.StatusBadGateway)
			return
		}

		sendJson(w, Response{Data: res}, http.StatusOK)
	}
}
