package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, status int, message string, data ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, message, ApiResponse{
		Status:  "error",
		Message: message,
	})
}

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	JSON(w, status, message, ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}
