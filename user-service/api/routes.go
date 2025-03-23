package api

import (
	"net/http"
	"user-service/api/handlers"
)

func SetupRoutes(handler *handlers.UserHandler, jwtsecret string) *http.ServeMux{ 
	mux:= http.NewServeMux()
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	return mux
}