package controller

import (
	"belajar-go/dto"
	"belajar-go/models"
	"belajar-go/services"
	"belajar-go/utils"
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

var validate = validator.New()

func (s *UserController) FindByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok := ctx.Value("user_id").(int)
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := s.service.GetProfile(ctx, userId)

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

	err = validate.Struct(InputUser)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
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

	token, err := s.service.Login(ctx, InputLogin.Email, InputLogin.Password)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "login successful", map[string]interface{}{"token": token})
}

func (s *UserController) TopUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId,ok := ctx.Value("user_id").(int)
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}
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

	err = s.service.TopUpSaldo(ctx, userId, InputTopUp.Amount)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "saldo topped up successfully", nil)
}
