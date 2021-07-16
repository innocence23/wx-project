package model

import "time"

type Permission struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;NOT NULL" json:"name"`                 // 权限名
	Group      string    `gorm:"column:group;NOT NULL" json:"group"`                 // 权限名
	Url       string    `gorm:"column:url;NOT NULL" json:"url"`                   // 路径
	Method    string    `gorm:"column:method;default:GET;NOT NULL" json:"method"` // 方法名称
	Status    int       `gorm:"column:status;default:1;NOT NULL" json:"status"`   // 状态: 1:正常; 2禁用
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}
