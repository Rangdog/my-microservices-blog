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

type FavoriteHandler struct {
	service *service.FavoriteService
	validator *validator.Validate
}

func NewFavoriteHandler(service *service.FavoriteService) *FavoriteHandler{
	return &FavoriteHandler{service: service, validator: validator.New()}
}

func (h *FavoriteHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		StoryID int64 `json:"storyId" validate:"required"`
		UserID int64 `json:"userId" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.CreateFavorite(input.StoryID, input.UserID)
	if err != nil{
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusCreated,"", "Create Favorite successfully")
}

func (h *FavoriteHandler) GetALLFavoriteByStoryID(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10 ,60)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	favorites,err := h.service.GetALLFavoriteByStoryID(id)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	data:=map[string]interface{}{
		"favorites": favorites,
	}
	response.Success(w, http.StatusOK, data, "successfuly")
}

func (h *FavoriteHandler) GetALLFavoriteByUserID(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10 ,60)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	favorites,err := h.service.GetALLFavoriteByUserID(id)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	data:=map[string]interface{}{
		"favorites": favorites,
	}
	response.Success(w, http.StatusOK, data, "successfuly")
}

func (h *FavoriteHandler) DeleteById(w http.ResponseWriter, r *http.Request){
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
	response.Success(w, http.StatusOK,"","delete successfuly")
}
