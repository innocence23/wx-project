package main

import (
	"log"
	"wx/app/handler"
	"wx/app/repository"
	"wx/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// will initialize a handler starting from data sources
// which inject into repository layer
// which inject into service layer
// which inject into handler layer
func inject(d *gorm.DB) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	// repository layer
	userRepository := repository.NewUserRepository(d)

	// repository layer
	userService := service.NewUserService(userRepository)

	// initialize gin.Engine
	router := gin.Default()
	handler.NewHandler(&handler.Config{
		R:           router,
		UserService: userService,
	})

	return router, nil
}
