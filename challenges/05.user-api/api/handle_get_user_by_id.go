package api

import (
	"errors"
	"net/http"

	"github.com/felipeolliveira/user_api_golang_challenge/database"
	"github.com/go-chi/chi/v5"
)

func handleGetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	user, err := database.FindById(userId)

	if err != nil {
		if errors.Is(err, database.ErrNotFoundUserFromId) {
			sendJson(
				w,
				Response{
					Error: err.Error(),
				},
				http.StatusNotFound,
			)

			return
		}

		sendJson(
			w,
			Response{
				Error: err.Error(),
			},
			http.StatusInternalServerError,
		)

		return
	}

	sendJson(
		w,
		Response{
			Data: user,
		},
		http.StatusOK,
	)
}
