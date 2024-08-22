package api

import (
	"errors"
	"net/http"

	"github.com/felipeolliveira/user_api_golang_challenge/database"
)

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.FindAll()

	if err != nil {
		if errors.Is(err, database.ErrCouldNotRetrieveUsers) {
			sendJson(
				w,
				Response{
					Error: err.Error(),
				},
				http.StatusInternalServerError,
			)
		}
		return
	}

	sendJson(
		w,
		Response{
			Data: users,
		},
		http.StatusOK,
	)
}
