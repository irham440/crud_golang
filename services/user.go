package services

import (
	"belajar-go/models"
	"belajar-go/storage"
	"belajar-go/utils"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserService interface{
	CreateUser(ctx context.Context, user *models.User) error
	GetProfile(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, email string, password string) (string, error)
	TopUpSaldo(ctx context.Context, id int, amount float64) error
}



type IUserService struct {
	repo storage.UserRepository
	db   *gorm.DB
}

func NewIUserService(repo storage.UserRepository, db *gorm.DB) *IUserService {
	return &IUserService{repo: repo, db: db}
}

func (s *IUserService) CreateUser(ctx context.Context, user *models.User) error {
	password, err := utils.GeneratePassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	fmt.Println(password)
	user.Password = password
	err = utils.ComparePassword([]byte(user.Password), []byte(password))
	fmt.Println("Hasil Compare Manual:", err == nil)
	err = s.repo.Save(ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}
func (s *IUserService) Login(ctx context.Context, email string, password string) (string ,error) {
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

func (s *IUserService) GetProfile(ctx context.Context, id int) (*models.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *IUserService) TopUpSaldo(ctx context.Context, id int, amount float64) error {
	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}
	err = s.db.Transaction(func(tx *gorm.DB) error {
		err = s.repo.AddSaldo(ctx, tx, id, amount)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to top up saldo")
	}
	return nil
}
