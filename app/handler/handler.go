package handler

import (
	"wx/app/handler/middleware"
	"wx/app/model"

	"github.com/gin-gonic/gin"
)

type Config struct {
	R           *gin.Engine
	BaseUrlPath string
	UserService model.UserService
	RoleService model.RoleService
}

func NewHandler(c *Config) {
	g := c.R.Group(c.BaseUrlPath)

	NewUserHandler(c.UserService).Router(g)

	g.Use(middleware.JWTAuthMiddleware())
	NewRoleHandler(c.RoleService).Router(g)
}
