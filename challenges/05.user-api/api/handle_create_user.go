package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/felipeolliveira/user_api_golang_challenge/api/validations"
	"github.com/felipeolliveira/user_api_golang_challenge/database"
)

type createUserPostResponse struct {
	Id string `json:"id"`
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
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

	err := validations.UserSchemaValidateNamesAndBio(body)
	var validationError validations.ErrUserSchema

	if err != nil {
		if errors.As(err, &validationError) {
			sendJson(
				w,
				Response{
					Error: validationError.Message,
					Data:  validationError.Fields,
				},
				http.StatusBadRequest,
			)
			return
		}
		slog.Error("Uncaught erro on validation:", "err", err)
		sendJson(
			w,
			Response{
				Error: "Something went wrong",
			},
			http.StatusInternalServerError,
		)
		return
	}

	createdUserId, err := database.Insert(body.FirstName,
		body.LastName,
		body.Bio)

	if err != nil {
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
			Data: createUserPostResponse{
				Id: createdUserId.String(),
			},
		},
		http.StatusCreated,
	)
}
