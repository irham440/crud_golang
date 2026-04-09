package user

import (
	"belajar-go/models"
	"belajar-go/services/email"
	"belajar-go/storage"
	"context"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetProfile(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, email string, password string) (string, error)
}

type IUserService struct {
	repo  storage.UserRepository
	db    *gorm.DB
	email email.EmailService
}

func NewIUserService(repo storage.UserRepository, db *gorm.DB, email email.EmailService) *IUserService {
	return &IUserService{repo: repo, db: db, email: email}
}
