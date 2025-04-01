package api

import (
	"user-service/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(handler *handlers.UserHandler, jwtsecret string) *chi.Mux{ 
	r:= chi.NewRouter()
	// Tạo một nhóm route với tiền tố `/api/user-service`
	r.Route("/api/user-service", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})
	r.Get("/health", handler.HealthCheck)
	return r
}