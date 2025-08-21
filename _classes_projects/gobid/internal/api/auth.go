package api

import (
	"net/http"
	"os"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/jsonutils"
	"github.com/gorilla/csrf"
)

const (
	SessionUserKey = "AuthenticatedUserId"
)

func (api *Api) handleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("GOBID_CSRF_KEY")

	if key == "" {
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, "CSRF key is required")
		return
	}

	token := csrf.Token(r)
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]string{"csrfToken": token})
}

func (api *Api) csrfMiddleware(next http.Handler) http.Handler {
	key := os.Getenv("GOBID_CSRF_KEY")
	isSecure := os.Getenv("GOBID_CSRF_SECURE") == "true"

	if key == "" {
		panic("CSRF key is required in production environment")
	}

	middleware := csrf.Protect([]byte(key), csrf.Secure(isSecure))
	return middleware(next)
}

func (api *Api) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Session.Exists(r.Context(), SessionUserKey) {
			jsonutils.EncodeJsonWithError(w, r, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
