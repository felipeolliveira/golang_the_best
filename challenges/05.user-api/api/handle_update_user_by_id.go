package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/felipeolliveira/user_api_golang_challenge/api/validations"
	"github.com/felipeolliveira/user_api_golang_challenge/database"
	"github.com/go-chi/chi/v5"
)

func handleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body validations.UserSchemaBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sendJson(
			w,
			Response{
				Error: "Please provide FirstName LastName and bio for the user",
			},
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()

	var validationErr validations.ErrUserSchema
	if err := validations.UserSchemaValidateNamesAndBio(body); err != nil {
		if errors.As(err, &validationErr) {
			sendJson(
				w,
				Response{
					Error: validationErr.Message,
					Data:  validationErr.Fields,
				},
				http.StatusBadRequest,
			)
			return
		}
		return
	}

	if err := database.Update(id, body.FirstName, body.LastName, body.Bio); err != nil {
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

	w.WriteHeader(http.StatusNoContent)
}
