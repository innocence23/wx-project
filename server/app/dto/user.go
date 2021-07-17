package dto

import (
	"wx/app/model"
)

type SignupReq struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required,gte=6,lte=30" label:"密码"`
}

type SigninReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

type DetailsReq struct {
	Account string `json:"account" binding:"omitempty,max=50"`
	Avatar  string `json:"avatar" binding:"required"`
}

type UserJWT struct {
	ID          int64               `json:"id"`
	Account     string              `json:"account"`
	Description string              `json:"description,omitempty"`
	Email       string              `json:"email"`
	Avatar      string              `json:"avatar"`
	Roles       []string            `json:"roles"`
	Menus       []model.Menu        `json:"menus"`
	Permissions []map[string]string `json:"permissions"`
}
