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

	// service
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)

	// handler
	router := gin.Default()
	handler.NewHandler(&handler.Config{
		R:           router,
		BaseUrlPath: os.Getenv("WX_API_URL_V1"),
		UserService: userService,
		RoleService: roleService,
	})

	return router, nil
}
