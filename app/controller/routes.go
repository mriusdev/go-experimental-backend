package controller

import (
	"api-backend/sample/app/application"
	"api-backend/sample/app/controller/middleware"

	"github.com/go-chi/chi/v5"
)

func AddRoutes(apiHandler *application.ApiHandler) chi.Router {
	mux := apiHandler.Mux
	authController := NewAuthController(apiHandler.Logger, apiHandler.DB)
	mux.Route("/register", func(r chi.Router) {
		r.Post("/", authController.Register)
	})
	mux.Route("/login", func(r chi.Router) {
		r.Post("/", authController.Login)
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.JwtAuthCheck)

		r.Route("/logout", func(r chi.Router) {
			r.Post("/", authController.Logout)
		})
		r.Route("/refresh", func(r chi.Router) {
			r.Post("/", authController.RefreshToken)
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", GetUsers)
			r.Post("/", CreateUser)
			r.Patch("/{id}", UpdateUser)
			r.Get("/{id}", GetUser)
			r.Delete("/{id}", DeleteUser)
		})
	})

	mux.Route("/info", func(r chi.Router) {
		r.Get("/", GetInfo)
	})

	return mux
}
