package component

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_ = db
}
