package handler

import (
	"wx/app/model"

	"github.com/gin-gonic/gin"
)

type Config struct {
	R           *gin.Engine
	UserService model.UserService
	BaseUrlPath string
}

func NewHandler(c *Config) {
	g := c.R.Group(c.BaseUrlPath)

	NewUserHandler(c.UserService).Router(g)
}
