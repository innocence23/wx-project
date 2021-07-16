package main

import (
	"log"
	"os"
	"wx/app/handler"
	"wx/app/repository"
	"wx/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func inject(d *gorm.DB) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	// repository
	userRepository := repository.NewUserRepository(d)
	roleRepository := repository.NewRoleRepository(d)
	permissionRepository := repository.NewPermissionRepository(d)
	menuRepository := repository.NewMenuRepository(d)

	// service
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)
	permissionService := service.NewPermissionService(permissionRepository)
	menuService := service.NewMenuService(menuRepository)

	// handler
	router := gin.Default()
	handler.NewHandler(&handler.Config{
		R:                 router,
		BaseUrlPath:       os.Getenv("APP_API_V1"),
		UserService:       userService,
		RoleService:       roleService,
		PermissionService: permissionService,
		MenuService:       menuService,
	})

	return router, nil
}
