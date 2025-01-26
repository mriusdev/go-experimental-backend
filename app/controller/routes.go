package controller

import (
	"github.com/go-chi/chi/v5"
)

func GetRouter() chi.Router {
	router := chi.NewRouter()
	router.Route("/login", func(r chi.Router) {
		r.Post("/", Login)
	})
	router.Route("/info", func(r chi.Router) {
		r.Get("/", GetInfo)
	})
	router.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsers)
		r.Post("/", CreateUser)
		r.Patch("/{id}", UpdateUser)
		r.Get("/{id}", GetUser)
		r.Delete("/{id}", DeleteUser)
	})

	return router
}
