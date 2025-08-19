package api

import (
	"taskify/internal/services"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	Router      *chi.Mux
	TaskService services.TaskService
}
