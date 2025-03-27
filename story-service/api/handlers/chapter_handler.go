package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"stories-service/internal/domain/service"
	response "stories-service/internal/pkg/Response"

	"github.com/go-playground/validator/v10"
)

type ChapterHandler struct {
	service *service.ChapterService
	validator *validator.Validate
}

func NewChapterHandler (service *service.ChapterService) *ChapterHandler{
	return &ChapterHandler{service: service, validator: validator.New()}
}


func (h *ChapterHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w,http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		Story_id int64 `json:"storyId" validate:"required"`
		Title string `json:"title"`
	}

	if err:=json.NewDecoder(r.Body).Decode(&input); err!=nil{
		response.Error(w, http.StatusBadRequest, err)
	}

	if err:= h.validator.Struct(input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	err:=h.service.CreateChapter(input.Story_id, input.Title)
	if err!=nil{
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusCreated, "",  "Create Chapter successfully")
}
