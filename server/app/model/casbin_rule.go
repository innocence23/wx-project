package model

import "time"

type CasbinRule struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	PType     string    `gorm:"column:p_type" json:"p_type"`
	V0        string    `gorm:"column:v0" json:"v0"`
	V1        string    `gorm:"column:v1" json:"v1"`
	V2        string    `gorm:"column:v2" json:"v2"`
	V3        string    `gorm:"column:v3" json:"v3"`
	V4        string    `gorm:"column:v4" json:"v4"`
	V5        string    `gorm:"column:v5" json:"v5"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
