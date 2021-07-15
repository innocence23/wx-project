package iface

import (
	"context"
	"wx/app/dto"
	"wx/app/model"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Signup(ctx context.Context, u *model.User) error
	Signin(ctx context.Context, u *model.User) error
	UpdateDetail(ctx context.Context, u *model.User) error
	DisableUser(ctx context.Context, id int64) error
	EnableUser(ctx context.Context, id int64) error
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, u *model.User) error
	Update(ctx context.Context, u *model.User) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}

type RoleService interface {
	Get(ctx context.Context, id int64) (*model.Role, error)
	List(ctx context.Context, where dto.RoleSearchReq) (dto.RoleListResp, error)
	Create(ctx context.Context, m *model.Role) (*model.Role, error)
	Update(ctx context.Context, m *model.Role) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
}

type RoleRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Role, error)
	FindByWhere(ctx context.Context, where dto.RoleSearchReq) (dto.RoleListResp, error)
	Create(ctx context.Context, m *model.Role) (*model.Role, error)
	Update(ctx context.Context, m *model.Role) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}

type PermissionService interface {
	Get(ctx context.Context, id int64) (*model.Permission, error)
	List(ctx context.Context, where dto.PermissionSearchReq) (dto.PermissionListResp, error)
	Create(ctx context.Context, m *model.Permission) (*model.Permission, error)
	Update(ctx context.Context, m *model.Permission) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
}

type PermissionRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Permission, error)
	FindByWhere(ctx context.Context, where dto.PermissionSearchReq) (dto.PermissionListResp, error)
	Create(ctx context.Context, m *model.Permission) (*model.Permission, error)
	Update(ctx context.Context, m *model.Permission) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}
