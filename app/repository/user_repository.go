package repository

import (
	"context"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() model.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Find(user, uid); err != nil {
		return user, zerror.NewNotFound("uid", uid.String())
	}
	return user, nil
}
