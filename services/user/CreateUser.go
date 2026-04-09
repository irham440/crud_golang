package user

import (
	"belajar-go/models"
	"belajar-go/utils"
	"context"
	"errors"
	"fmt"
	"time"
)

func (s *IUserService) CreateUser(ctx context.Context, user *models.User) error {
	password, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = password
	err = s.repo.Save(ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}
	body := fmt.Sprintf("anda berhasil buat akun pada %s", time.Now().Format("01 Jan 2006 15.04"))
	s.email.SendEmail(user.Email, "daftar akun", body)

	return nil
}
