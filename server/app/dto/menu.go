package dto

import (
	"time"
	"wx/app/model"
)

type MenuIdReq struct {
	ID int64 `json:"id" form:"id" binding:"required" label:"ID"`
}

type MenuCreateReq struct {
	PId         *int64 `json:"p_id" binding:"required"` // 可能为0
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
	Weight      int    `json:"weight" binding:"required"`
}

type MenuUpdateReq struct {
	ID          int64  `json:"id" binding:"required"`
	PId         *int64 `json:"p_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Url         string `json:"url" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
	Weight      int    `json:"weight" binding:"required"`
}

type MenuSearchReq struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	CreatedAtMin string `json:"created_at_min"`
	CreatedAtMax string `json:"created_at_max"`
	Pagination
}

type MenuListResp struct {
	List  []model.Menu `json:"list"`
	Total int64        `json:"total"`
	Pagination
}

type Menu struct {
	Id int64 `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`

	Status    int       `gorm:"column:status;default:1;NOT NULL" json:"status"` // 状态: 1: 启用; 2:禁用;
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}
