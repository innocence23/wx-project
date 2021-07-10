package service

import (
	"context"
	"wx/app/model"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository model.UserRepository
}

func NewUserService(ur model.UserRepository) model.UserService {
	return &UserService{
		UserRepository: ur,
	}
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}
