package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/internal/domain/service"
	response "user-service/internal/pkg/Response"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service *service.UserService
	validator *validator.Validate
}

func NewUserHandler(service *service.UserService) *UserHandler{
	return &UserHandler{service: service, validator: validator.New()}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request){
	var input struct{
		Email string `json:"email" validate:"required, email"`
		Password string `json:"password" validate:"required, min 6"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
	}

	user, err := h.service.Register(input.Email, input.Password)
	if err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}
	response.Success(w, http.StatusCreated,user, "User registered successfully")
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request){
	var input struct{
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	if err:=json.NewDecoder(r.Body).Decode(&input); err != nil{
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err:=h.validator.Struct(input); err!=nil{
		response.Error(w,http.StatusBadRequest, err)
		return
	}

	user,token, err := h.service.Login(input.Email, input.Password)
	if err != nil{
		response.Error(w, http.StatusUnauthorized,err)
		return
	}

	data:=map[string]interface{}{
		"user": user,
		"token": token,
	}
	response.Success(w, http.StatusOK, data, "Login successfuly")
}

func (h *UserHandler) HealthCheck(w http.ResponseWriter, r *http.Request){
	response:=map[string]string{"status": "health"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}