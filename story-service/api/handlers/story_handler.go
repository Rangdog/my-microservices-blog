package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"stories-service/internal/domain/service"
	response "stories-service/internal/pkg/Response"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type StoryHandler struct {
	service *service.StoryService
	validator *validator.Validate
}

func NewStoryHandler(service *service.StoryService) *StoryHandler{
	return &StoryHandler{service: service, validator: validator.New()}
}

func (h *StoryHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		Title string `json:"title" validate:"required"`
		Description string `json:"description"`
		Author_id int64 `json:"author_id" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	err := h.service.CreateStory(input.Title, input.Description, input.Author_id)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Success(w, http.StatusCreated,"", "Create Story successfully")
}

func (h *StoryHandler) FindById(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	idStr := chi.URLParam(r,"id")
	id, err := strconv.ParseInt(idStr,10,64)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	story, err := h.service.GetStoryByID(id)
	if err != nil{
		response.Error(w, http.StatusBadRequest,err)
		return
	}

	data:=map[string]interface{}{
		"story": story,
	}
	response.Success(w, http.StatusOK, data, "successfuly")
}

func (h *StoryHandler) DeleteById(w http.ResponseWriter, r *http.Request){
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

func (h *StoryHandler) HealthCheck(w http.ResponseWriter, r *http.Request){
	response:=map[string]string{"status": "health"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}