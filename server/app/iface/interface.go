package iface

import (
	"context"
	"wx/app/dto"
	"wx/app/model"
)

type UserService interface {
	Find(ctx context.Context, id int64) (*model.User, error)
	List(ctx context.Context, where dto.UserSearchReq) (dto.UserListResp, error)
	Signup(ctx context.Context, u *model.User) error
	Signin(ctx context.Context, u *model.User) error
	Update(ctx context.Context, u *model.User) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
	Resetpwd(ctx context.Context, id int64) error
	GetMenus(ctx context.Context, roles []int64, email string) []model.Menu
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByWhere(ctx context.Context, where dto.UserSearchReq) (dto.UserListResp, error)
	Create(ctx context.Context, u *model.User) error
	Update(ctx context.Context, u *model.User) error
	UpdatePassword(ctx context.Context, u *model.User) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}

type RoleService interface {
	Find(ctx context.Context, id int64) (*model.Role, error)
	List(ctx context.Context, where dto.RoleSearchReq) (dto.RoleListResp, error)
	Create(ctx context.Context, m *model.Role) (*model.Role, error)
	Update(ctx context.Context, m *model.Role) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
}

type RoleRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Role, error)
	FindByIds(ctx context.Context, ids []int64) ([]model.Role, error)
	FindByWhere(ctx context.Context, where dto.RoleSearchReq) (dto.RoleListResp, error)
	Create(ctx context.Context, m *model.Role) (*model.Role, error)
	Update(ctx context.Context, m *model.Role) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}

type PermissionService interface {
	Find(ctx context.Context, id int64) (*model.Permission, error)
	List(ctx context.Context, where dto.PermissionSearchReq) (dto.PermissionListResp, error)
	Create(ctx context.Context, m *model.Permission) (*model.Permission, error)
	Update(ctx context.Context, m *model.Permission) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
	AutoGenerate(ctx context.Context, routers []map[string]string) error
}

type PermissionRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Permission, error)
	FindByWhere(ctx context.Context, where dto.PermissionSearchReq) (dto.PermissionListResp, error)
	Create(ctx context.Context, m *model.Permission) (*model.Permission, error)
	Update(ctx context.Context, m *model.Permission) error
	UpdateStatus(ctx context.Context, id int64, status int) error
	FindByUrlAndMethod(ctx context.Context, url, method string) (*model.Permission, error)
}

type MenuService interface {
	Find(ctx context.Context, id int64) (*model.Menu, error)
	List(ctx context.Context, where dto.MenuSearchReq) (dto.MenuListResp, error)
	Create(ctx context.Context, m *model.Menu) (*model.Menu, error)
	Update(ctx context.Context, m *model.Menu) error
	Disable(ctx context.Context, id int64) error
	Enable(ctx context.Context, id int64) error
}

type MenuRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Menu, error)
	FindByIds(ctx context.Context, ids []int64) ([]model.Menu, error)
	FindAll(ctx context.Context) ([]model.Menu, error)
	FindByWhere(ctx context.Context, where dto.MenuSearchReq) (dto.MenuListResp, error)
	Create(ctx context.Context, m *model.Menu) (*model.Menu, error)
	Update(ctx context.Context, m *model.Menu) error
	UpdateStatus(ctx context.Context, id int64, status int) error
}
