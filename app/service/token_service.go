package service

import (
	"context"
	"wx/app/model"
)

type TokenService struct {
}

func NewTokenService() model.TokenService {
	return &TokenService{}
}

func (s *TokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	panic("Method not implemented")
}
