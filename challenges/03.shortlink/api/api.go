package api

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleResolveCode(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJson(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJson(w, Response{Error: "invalid url passed"}, http.StatusBadRequest)
		}

		code := genCode()
		db[code] = body.URL

		type codeResponse struct {
			Code string `json:"code"`
		}
		sendJson(w, Response{Data: codeResponse{Code: code}}, http.StatusCreated)
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func genCode() string {
	const n = 8
	codeBytes := make([]byte, n)
	for i := range n {
		codeBytes[i] = characters[rand.IntN(len(characters))]
	}
	return string(codeBytes)
}

func handleResolveCode(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		foundUrl, ok := db[code]

		if !ok {
			http.Error(w, "Short link não encontrado", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, foundUrl, http.StatusPermanentRedirect)
	}
}
