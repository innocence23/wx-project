package model

import "time"

type User struct {
	ID        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;NOT NULL" json:"name"`
	Email     string    `gorm:"column:email;NOT NULL" json:"email"`
	Password  string    `gorm:"column:password;NOT NULL" json:"password"`
	Avatar    string    `gorm:"column:avatar;NOT NULL" json:"avatar"`
	Status    int       `gorm:"column:status;default:1;NOT NULL" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}
