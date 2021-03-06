package dto

import "wx/app/model"

type PermissionIdReq struct {
	ID int64 `json:"id" form:"id" binding:"required" label:"ID"`
}

type PermissionCreateReq struct {
	Name   string `json:"name" binding:"required"`
	Group  string `json:"group" binding:"required"`
	Url    string `json:"url" binding:"required"`
	Method string `json:"method" binding:"required"`
}

type PermissionUpdateReq struct {
	ID     int64  `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Group  string `json:"group" binding:"required"`
	Url    string `json:"url" binding:"required"`
	Method string `json:"method" binding:"required"`
}

type PermissionSearchReq struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	CreatedAtMin string `json:"created_at_min"`
	CreatedAtMax string `json:"created_at_max"`
	Pagination
}

type PermissionListResp struct {
	List  []model.Permission `json:"list"`
	Total int64              `json:"total"`
	Pagination
}
