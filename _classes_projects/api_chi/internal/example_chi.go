package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ChiExample() *chi.Mux {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/now", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, now)
	})

	r.Route("/api", func(r chi.Router) {
		// Route agrupa as rotas passando subrotas
		r.Route("/v1", func(r chi.Router) {
			// Pode usar os parametros de rota igual a standard lib
			r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				fmt.Fprintln(w, "users", id)
			})
			// Mas também é possível usar regexp
			r.Get("/products/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				fmt.Fprintln(w, "products", id)
			})
		})
		r.Route("/v2", func(r chi.Router) {
			r.Get("/users", func(w http.ResponseWriter, r *http.Request) {})
		})

		// With serve para colocar um middleware antes de uma definição de rota
		r.With(middleware.RealIP).Get("/users", func(w http.ResponseWriter, r *http.Request) {})

		// Group serve para agrupar rotas sem criar subrotas, serve para usar middlewares isolados dentro de grupos específicos
		r.Group(func(r chi.Router) {
			// Assim, o Use serve como middleware apenas para essas rotas dentro desse Group
			r.Use(middleware.BasicAuth("", map[string]string{
				"admin": "admin",
			}))

			r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Authenticated route")
			})
		})
	})

	return r
}
