package controller

import (
	"belajar-go/dto"
	"belajar-go/models"
	"belajar-go/services"
	"belajar-go/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (s *UserController) FindByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid ID")
		return
	}

	user, err := s.service.GetProfile(ctx, id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "user profile retrieved successfully", user)
}

func (s *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var InputUser dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&InputUser)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user := models.User{
		Name:     InputUser.Name,
		Email:    InputUser.Email,
		Password: InputUser.Password,
	}
	err = s.service.CreateUser(ctx, &user)
	if err != nil {
		utils.Error(w, http.StatusConflict, err.Error())
		return
	}

	utils.Success(w, http.StatusCreated, "user created successfully", user)
}

func (s *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var InputLogin dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&InputLogin)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = s.service.Login(ctx, InputLogin.Email, InputLogin.Password)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	utils.Success(w, http.StatusOK, "login successful", nil)
}

func (s *UserController) TopUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var InputTopUp dto.TopUpSaldoRequest
	err := json.NewDecoder(r.Body).Decode(&InputTopUp)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if InputTopUp.Amount <= 0 {
		utils.Error(w, http.StatusBadRequest, "amount must be greater than zero")
		return
	}

	err = s.service.TopUpSaldo(ctx, InputTopUp.Id, InputTopUp.Amount)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "saldo topped up successfully", nil)
}
