package dto

import "wx/app/model"

type RoleIdReq struct {
	ID int64 `json:"id" form:"id" binding:"required" label:"ID"`
}

type RoleCreateReq struct {
	Name          string        `json:"name" binding:"required"`
	Description   string        `json:"description" binding:"required"`
	MenuIds       model.IntJson `json:"menu_ids" binding:"required"`
	PermissionIds []int         `json:"permission_ids"`
}

type RoleUpdateReq struct {
	ID            int64         `json:"id" binding:"required"`
	Name          string        `json:"name" binding:"required"`
	Description   string        `json:"description" binding:"required"`
	MenuIds       model.IntJson `json:"menu_ids" binding:"required"`
	PermissionIds []int         `json:"permission_ids"`
}

type RoleSearchReq struct {
	Name         string `json:"name"`
	CreatedAtMin string `json:"created_at_min"`
	CreatedAtMax string `json:"created_at_max"`
	Pagination
}

type RoleListResp struct {
	List  []model.Role `json:"list"`
	Total int64        `json:"total"`
	Pagination
}
