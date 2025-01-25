package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/jsonutils"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/services"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		jsonutils.EncodeJsonWithError(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(), data.UserName, data.Email, data.Password, data.Bio)
	if err != nil {
		if errors.Is(err, services.UserServiceErr.ErrEmailOrUserNameAlreadyExists) {
			jsonutils.EncodeJsonWithError(w, r, http.StatusUnprocessableEntity, err.Error())
			return
		}
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)
	if err != nil {
		jsonutils.EncodeJsonWithError(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	userId, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, services.UserServiceErr.ErrInvalidCredentials) {
			jsonutils.EncodeJsonWithError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	err = api.Session.RenewToken(r.Context())
	if err != nil {
		slog.Error("handleLoginUser", "error renewing token", err)
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, "internal server error")
	}

	api.Session.Put(r.Context(), SessionUserKey, userId)
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged in",
	})
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Session.RenewToken(r.Context())
	if err != nil {
		slog.Error("handleLogoutUser", "error renewing token", err)
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, "internal server error")
	}

	api.Session.Remove(r.Context(), SessionUserKey)
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "logged out",
	})
}
