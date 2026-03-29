package services

import (
	"belajar-go/models"
	"belajar-go/storage"
)

type UserService struct {
	repo storage.UserRepository
}

func NewUserService(repo storage.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindById(id int) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Create(user models.User) error {
	return s.repo.Save(user)
}

func (s *UserService) FindEmail(email string) (string, error) {
	return s.repo.GetEmail(email)
}	