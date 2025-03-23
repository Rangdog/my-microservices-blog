package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json="status"`
	Message string      `json="message,omitempty"`
	Data    interface{} `json="data,omitempty"`
	Error   string      `json="error,omitempty"`
}

func Success(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Status: "success",
		Message: message,
		Data: data,
	})
}

func Error(w http.ResponseWriter, statusCode int, err error){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Status: "error",
		Error: err.Error(),
	})
}