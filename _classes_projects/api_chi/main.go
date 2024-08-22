package main

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/felipeolliveira/api_chi_golang/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

type User struct {
	Id       int64  `json:"id,string"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"-"`
}

func sendJson(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Erro ao fazer marshal de json", "error", err)
		sendJson(w, Response{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Erro ao enviar a resposta", "error", err)
		return
	}
}

func main() {
	// Exemplo de Chi, para consultar em caso de dúvida
	_ = internal.ChiExample()

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	inMemoryDB := map[int64]User{
		1: {
			Username: "admin",
			Id:       1,
			Role:     "admin",
			Password: "admin",
		},
	}

	r.Group(func(r chi.Router) {
		r.Use(internal.JsonMiddleware)

		r.Post("/users", handlePostUsers(inMemoryDB))
		r.Get("/users/{id:[0-9]+}", handleGetUsers(inMemoryDB))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawId := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(rawId, 10, 64)

		user, ok := db[id]
		if !ok {
			sendJson(w, Response{Error: "usuário não encontrado"}, http.StatusNotFound)
			return
		}

		sendJson(
			w,
			Response{Data: user},
			http.StatusOK,
		)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10000)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var errMaxByte *http.MaxBytesError
			if errors.As(err, &errMaxByte) {
				sendJson(
					w,
					Response{Error: "body too large"},
					http.StatusRequestEntityTooLarge,
				)
				return
			}
			slog.Error("Falha ao ler o json do usuário", "error", err)
			sendJson(
				w,
				Response{Error: "Something went wrong"},
				http.StatusInternalServerError,
			)
			return
		}
		var user User
		if err := json.Unmarshal(data, &user); err != nil {
			sendJson(
				w,
				Response{Error: "invalid body"},
				http.StatusUnprocessableEntity,
			)
			return
		}

		user.Id = int64(len(db) + 1)
		db[user.Id] = user

		type idResponse struct {
			Id int64 `json:"id,string"`
		}
		sendJson(w, Response{Data: idResponse{Id: user.Id}}, http.StatusCreated)
	}
}
