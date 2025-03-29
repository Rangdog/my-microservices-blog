package api

import (
	"Interaction-service/api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(commentHandler *handlers.CommentHandler, favoriteHandler *handlers.FavoriteHandler, followHandler *handlers.FollowHandler, ratingHandler *handlers.RatingHandler) http.Handler{ 
	r:=chi.NewRouter()
	//comment
	r.Post("/comment", commentHandler.Create)
	r.Get("/comment/stories/{id}", commentHandler.GetALLCommentByStoryID)
	r.Delete("/comment",commentHandler.DeleteById)

	//favorite
	r.Post("/favotite", favoriteHandler.Create)
	r.Get("/favotite/stories/{id}", favoriteHandler.GetALLFavoriteByStoryID)
	r.Get("/favotite/users/{id}", favoriteHandler.GetALLFavoriteByUserID)
	r.Delete("/favorite",favoriteHandler.DeleteById)

	//follow
	r.Post("follow", followHandler.Create)
	r.Get("/follow/stories/{id}", followHandler.GetALLFolowByStoryID)
	r.Get("/follow/users/{id}", followHandler.GetALLFollowByUserID)
	r.Delete("/follow",favoriteHandler.DeleteById)

	//rating
	r.Post("/rating", ratingHandler.Create)
	r.Get("/rating/stories/{id}", ratingHandler.GetALLRattingByStoryID)
	r.Delete("/rating",ratingHandler.DeleteById)


	//health
	r.Get("/heath", commentHandler.HealthCheck)
	return r
}