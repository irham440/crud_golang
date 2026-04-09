package user

import (
	"belajar-go/models"
	"context"
	"errors"
)

func (s *IUserService) GetProfile(ctx context.Context, id int) (*models.User, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
