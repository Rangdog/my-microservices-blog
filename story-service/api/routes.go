package api

import (
	"net/http"
	"stories-service/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(storyHandler *handlers.StoryHandler, genreHandler *handlers.GenreHandler, chapterHandler *handlers.ChapterHandler, storyGenreHandler *handlers.StoryGenreHandler) http.Handler{ 
	r:=chi.NewRouter()
	r.Post("/stories", storyHandler.Create)
	r.Get("/stories/{id}", storyHandler.FindById)
	r.Delete("/stories/{id}", storyHandler.DeleteById)
	r.Post("/genre", genreHandler.Create)
	r.Post("/chapter", chapterHandler.Create)
	r.Post("/storygenre", storyGenreHandler.Create)
	r.Delete("/storygenre", storyGenreHandler.DeleteById)
	r.Get("/heath", storyHandler.HealthCheck)
	return r
}