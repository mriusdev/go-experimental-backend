package controller

import (
	"api-backend/sample/app/controller/middleware"

	"github.com/go-chi/chi/v5"
)

func GetRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/register", func(r chi.Router) {
		r.Post("/", Register)
	})
	router.Route("/login", func(r chi.Router) {
		r.Post("/", Login)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.JwtAuthCheck)

		r.Route("/logout", func(r chi.Router) {
			r.Post("/", Logout)
		})
		r.Route("/refresh", func(r chi.Router) {
			r.Post("/", RefreshToken)
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", GetUsers)
			r.Post("/", CreateUser)
			r.Patch("/{id}", UpdateUser)
			r.Get("/{id}", GetUser)
			r.Delete("/{id}", DeleteUser)
		})
	})

	router.Route("/info", func(r chi.Router) {
		r.Get("/", GetInfo)
	})

	return router
}
