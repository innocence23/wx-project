package component

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var enforcer *casbin.Enforcer

func InitCasbin(db *gorm.DB) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal(err)
	}
	// 通过mysql适配器新建一个enforcer
	enforcer, err = casbin.NewEnforcer("config/casbin.conf", adapter)
	if err != nil {
		log.Fatal("=====", err)
	}
	// 日志记录
	enforcer.EnableLog(true)
}
