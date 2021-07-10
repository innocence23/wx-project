package repository

import (
	"context"
	"log"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Find(user, uid); err != nil {
		return user, zerror.NewNotFound("uid", uid.String())
	}
	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	if err := r.DB.Create(u); err != nil {
		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return zerror.NewInternal()
	}
	return nil
}
