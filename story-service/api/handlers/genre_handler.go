package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"stories-service/internal/domain/service"
	response "stories-service/internal/pkg/Response"

	"github.com/go-playground/validator/v10"
)

type GenreHandler struct {
	service *service.GenreService
	validator *validator.Validate
}

func NewGenreHandler (service *service.GenreService) *GenreHandler{
	return &GenreHandler{service: service, validator: validator.New()}
}
func (h *GenreHandler) Create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		response.Error(w,http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}
	var input struct{
		Name string `json:"name" validate:"required"`
	}

	if err:=json.NewDecoder(r.Body).Decode(&input); err!=nil{
		response.Error(w, http.StatusBadRequest, err)
	}

	if err:= h.validator.Struct(input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	err:=h.service.CreateGenre(input.Name)
	if err!=nil{
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusCreated, "",  "Create Genre successfully")
}
