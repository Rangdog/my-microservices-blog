package api

import (
	"user-service/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(handler *handlers.UserHandler, jwtsecret string) *chi.Mux{ 
	r:= chi.NewRouter()
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)
	r.Get("/heath", handler.HealthCheck)
	return r
}