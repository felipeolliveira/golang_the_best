package api

import (
	"errors"
	"net/http"

	"github.com/felipeolliveira/user_api_golang_challenge/database"
	"github.com/go-chi/chi/v5"
)

func handleRemoveUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := database.Delete(id); err != nil {
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
				Error: "Could not delete user. Something went wrong!",
			},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
