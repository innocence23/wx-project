package dto

import (
	"wx/app/model"
)

type UserIdReq struct {
	ID int64 `json:"id" form:"id" binding:"required" label:"ID"`
}

type SignupReq struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required,gte=6,lte=30" label:"密码"`
}

type SigninReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type UserUpdateReq struct {
	Account string `json:"account" binding:"required,max=50"`
	Avatar  string `json:"avatar" binding:"required"`
}

type UserAddRolesReq struct {
	ID      int64    `json:"id" form:"id" binding:"required" label:"ID"`
	RoleIds []string `json:"role_ids" form:"id" role_ids:"required"`
}

type UserSearchReq struct {
	Account      string `json:"account"`
	Email        string `json:"email"`
	CreatedAtMin string `json:"created_at_min"`
	CreatedAtMax string `json:"created_at_max"`
	Pagination
}

type UserListResp struct {
	List  []model.User `json:"list"`
	Total int64        `json:"total"`
	Pagination
}

type UserJWT struct {
	ID          int64               `json:"id"`
	Account     string              `json:"account"`
	Email       string              `json:"email"`
	Avatar      string              `json:"avatar"`
	Roles       []string            `json:"roles"` //ids
	Menus       []model.Menu        `json:"menus"`
	Permissions []map[string]string `json:"permissions"`
}
