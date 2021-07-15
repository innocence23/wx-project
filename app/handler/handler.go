package handler

import (
	"wx/app/handler/middleware"
	"wx/app/iface"

	"github.com/gin-gonic/gin"
)

type Config struct {
	R                 *gin.Engine
	BaseUrlPath       string
	UserService       iface.UserService
	RoleService       iface.RoleService
	PermissionService iface.PermissionService
}

func NewHandler(c *Config) {
	g := c.R.Group(c.BaseUrlPath)

	NewUserHandler(c.UserService).Router(g)

	g.Use(middleware.JWTAuthMiddleware())
	NewRoleHandler(c.RoleService).Router(g)
	NewPermissionHandler(c.PermissionService).Router(g)
}
