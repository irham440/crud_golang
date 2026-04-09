package handlerUser

import (
	"belajar-go/dto"
	"belajar-go/utils"
	"encoding/json"
	"net/http"
)

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
