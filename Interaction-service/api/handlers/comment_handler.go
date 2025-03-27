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

type CommentHandler struct {
	service *service.CommentService
	validator *validator.Validate
}

func NewCommentHandler(service *service.CommentService) *CommentHandler{
	return &CommentHandler{service: service, validator: validator.New()}
}

func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		StoryID int64 `json:"storyId" validate:"required"`
		UserID int64 `json:"userId" validate:"required"`
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

	err := h.service.CreateComment(input.StoryID, input.UserID, input.Content)
	if err != nil{
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusCreated,"", "Create Comment successfully")
}

func (h *CommentHandler) GetALLCommentByStoryID(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		response.Error(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10 ,60)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	comments,err := h.service.GetALLCommentByStoryID(id)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	data:=map[string]interface{}{
		"comments": comments,
	}
	response.Success(w, http.StatusOK, data, "successfuly")
}

func (h *CommentHandler) DeleteById(w http.ResponseWriter, r *http.Request){
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

func (h *CommentHandler) HealthCheck(w http.ResponseWriter, r *http.Request){
	response:=map[string]string{"status": "health"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}