package application

import (
	"api-backend/sample/app/service/auth"

	"github.com/go-chi/chi/v5"
)

type App struct {
	router chi.Router
	auth   *auth.Auth
}
