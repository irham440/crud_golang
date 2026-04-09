package handlerUser

import (
	"belajar-go/dto"
	"belajar-go/models"
	"belajar-go/utils"
	"encoding/json"
	"net/http"
)

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
