package model

import "time"

type Role struct {
	Id          int64   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name        string  `gorm:"column:name;NOT NULL" json:"name"`               // 角色名
	Description string  `gorm:"column:description;NOT NULL" json:"description"` // 角色描述
	MenuIds     IntJson `gorm:"column:menu_ids;NOT NULL" json:"menu_ids"`       // 拥有的菜单
	//MenuIds         []int64]   `gorm:"column:menu_ids;NOT NULL" json:"menu_ids"`       // 拥有的菜单
	Status          int       `gorm:"column:status;default:1;NOT NULL" json:"status"` // 状态: 1: 启用; 2:禁用;
	CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	Permissions_ids []int     `json:"permission_ids" gorm:"-"`
}
