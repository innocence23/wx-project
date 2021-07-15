package model

import (
	"context"
)

// UserService defines methods the handler layer expects
type UserService interface {
	Get(ctx context.Context, id int64) (*User, error)
	Signup(ctx context.Context, u *User) error
	Signin(ctx context.Context, u *User) error
	UpdateDetail(ctx context.Context, u *User) error
	DisableUser(ctx context.Context, id int64) error
	EnableUser(ctx context.Context, id int64) error
}

// UserRepository defines methods the service layer expects
type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}
