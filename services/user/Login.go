package user

import (
	"belajar-go/utils"
	"context"
	"errors"
	"fmt"
)

func (s *IUserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	fmt.Println(password)
	fmt.Println(user.Password)
	err = utils.ComparePassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid or password")
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
