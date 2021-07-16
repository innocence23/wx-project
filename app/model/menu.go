package model

import "time"

type Menu struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PId         int64     `gorm:"column:p_id;default:1;NOT NULL" json:"p_id"` // 父ID
	Name        string    `gorm:"column:name;NOT NULL" json:"name"`
	Description string    `gorm:"column:description;NOT NULL" json:"description"` // 描述
	Url         string    `gorm:"column:url;NOT NULL" json:"url"`
	Icon        string    `gorm:"column:icon;NOT NULL" json:"icon"`
	Weight      int       `gorm:"column:weight;default:1;NOT NULL" json:"weight"` // 排序
	Status      int       `gorm:"column:status;default:1;NOT NULL" json:"status"` // 状态: 1: 启用; 2:禁用;
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}
