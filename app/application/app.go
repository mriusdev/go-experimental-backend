package application

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type ApiHandler struct {
	Mux *chi.Mux
	Logger *slog.Logger
	DB *gorm.DB
}

func NewApiHandler(mux *chi.Mux, logger *slog.Logger, db *gorm.DB) *ApiHandler {
	return &ApiHandler{
		Mux: mux,
		Logger: logger,
		DB: db,
	}
}
