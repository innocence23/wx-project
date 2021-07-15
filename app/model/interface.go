package model

import (
	"context"
	"wx/app/dto"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*User, error)
	Signup(ctx context.Context, u *User) error
	Signin(ctx context.Context, u *User) error
	UpdateDetail(ctx context.Context, u *User) error
	DisableUser(ctx context.Context, id int64) error
	EnableUser(ctx context.Context, id int64) error
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}

type RoleService interface {
	Get(ctx context.Context, id int64) (*Role, error)
	List(ctx context.Context, where dto.RoleSearchReq) ([]Role, error)
	Create(ctx context.Context, m *Role) (*Role, error)
	Update(ctx context.Context, m *Role) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
}

type RoleRepository interface {
	FindByID(ctx context.Context, id int64) (*Role, error)
	FindByWhere(ctx context.Context, where dto.RoleSearchReq) ([]Role, error)
	Create(ctx context.Context, m *Role) (*Role, error)
	Update(ctx context.Context, m *Role) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}
