package api

import (
	"net/http"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/jsonutils"
)

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Session.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJsonWithError(w, r, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
