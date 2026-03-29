package controller

import (
	"belajar-go/models"
	"belajar-go/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (s *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := s.service.FindById(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (s *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var InputUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&InputUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user := models.User{
		Name:     InputUser.Name,
		Email:    InputUser.Email,
		Password: InputUser.Password,
	}
	exits, _ := s.service.FindEmail(InputUser.Email)
	fmt.Println(exits)
	if exits == InputUser.Email {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}
	err = s.service.Create(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
