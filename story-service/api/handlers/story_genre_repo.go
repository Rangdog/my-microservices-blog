package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"stories-service/internal/domain/service"
	response "stories-service/internal/pkg/Response"

	"github.com/go-playground/validator/v10"
)

type StoryGenreHandler struct {
	service *service.StoryGenreService
	validator *validator.Validate
}

func NewStoryGenreHandler(service *service.StoryGenreService) *StoryGenreHandler{
	return &StoryGenreHandler{service: service, validator: validator.New()}
}

func (h *StoryGenreHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		StoryID int64 `json:"storyId" validate:"required"`
		GenreID int64 `json:"genreId" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.CreateStoryGenre(input.StoryID, input.GenreID)
	if err != nil{
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusCreated,"", "Create Story Genre successfully")
}

func (h *StoryGenreHandler) DeleteById(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
	var input struct{
		StoryID int64 `json:"storyId" validate:"required"`
		GenreID int64 `json:"genreId" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.DeleteById(input.StoryID,input.GenreID)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
}
