package services

import (
	"belajar-go/models"
	"belajar-go/storage"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
	repo storage.UserRepository
	db   *gorm.DB
}

func NewUserService(repo storage.UserRepository, db *gorm.DB) *UserService {
	return &UserService{repo: repo, db: db}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	_, err := s.repo.GetEmail(ctx, user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	err = s.repo.Save(ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (s *UserService) Login(ctx context.Context, email string, password string) error {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}
	if user.Password != password {
		return errors.New("invalid password")
	}
	return nil
}

func (s *UserService) GetProfile(ctx context.Context, id int) (*models.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) TopUpSaldo(ctx context.Context, id int, amount float64) error {
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
