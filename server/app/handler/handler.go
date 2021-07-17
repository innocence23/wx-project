package handler

import (
	"wx/app/component"
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
	MenuService       iface.MenuService
}

func NewHandler(c *Config) {
	grouter := c.R.Group(c.BaseUrlPath)

	grouter.Use(middleware.UrlPathMiddleware())
	NewUserHandler(c.UserService).Router(grouter)

	grouter.Use(middleware.JWTAuthMiddleware())
	//grouter.Use(middleware.RbacMiddleware())
	NewAdminHandler(c.UserService).Router(grouter)
	NewRoleHandler(c.RoleService).Router(grouter)
	NewPermissionHandler(c.PermissionService).Router(grouter)
	NewMenuHandler(c.MenuService).Router(grouter)

	component.Router = c.R
}
