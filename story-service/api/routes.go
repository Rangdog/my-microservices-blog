package api

import (
	"net/http"
	"stories-service/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(handler *handlers.StoryHandler, jwtsecret string) http.Handler{ 
	r:=chi.NewRouter()
	r.Post("/stories", handler.Create)
	r.Get("/stories/{id}", handler.FindById)
	r.Delete("/stories/{id}", handler.DeleteById)
	r.Get("/heath", handler.HealthCheck)
	return r
}