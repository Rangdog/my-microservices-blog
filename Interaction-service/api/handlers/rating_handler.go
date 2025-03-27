package handlers

import (
	"Interaction-service/internal/domain/service"
	response "Interaction-service/internal/pkg/Response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type RatingHandler struct {
	service *service.RatingService
	validator *validator.Validate
}

func NewRatingHandler(service *service.RatingService) *RatingHandler{
	return &RatingHandler{service: service, validator: validator.New()}
}

func (h *RatingHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		StoryID int64 `json:"storyId" validate:"required"`
		UserID int64 `json:"userId" validate:"required"`
		Rating int64 `json:"rating" validate:"required"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.CreateRatting(input.StoryID, input.UserID, input.Rating, input.Content)
	if err != nil{
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusCreated,"", "Create ratting successfully")
}

func (h *RatingHandler) GetALLRattingByStoryID(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10 ,60)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	ratings,err := h.service.GetALLRattingByStoryID(id)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	data:=map[string]interface{}{
		"ratings": ratings,
	}
	response.Success(w, http.StatusOK, data, "successfuly")
}


func (h *RatingHandler) DeleteById(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10 ,60)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DeleteById(id)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
}
