package handlerUser

import (
	"belajar-go/services/user"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service user.UserService
}

func NewUserController(service user.UserService) *UserController {
	return &UserController{service: service}
}

var validate = validator.New()
