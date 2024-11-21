package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router         *chi.Mux
	UserService    services.UserService
	ProductService services.ProductService
	Session        *scs.SessionManager
}
