package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)
	api.Router.Use(api.Session.LoadAndSave, api.csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/csrf-token", api.handleGetCSRFToken)

			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)

				r.Group(func(r chi.Router) {
					r.Use(api.authMiddleware)
					r.Post("/logout", api.handleLogoutUser)
				})
			})

			r.Route("/products", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(api.authMiddleware)
					r.Post("/", api.handleCreateProduct)

					r.Route("/ws", func(r chi.Router) {
						r.Get("/subscribe/{productId}", api.handleSubscribeUserToAuction)
					})
				})
			})
		})
	})
}
