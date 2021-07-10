package handler

import (
	"os"
	"wx/app/model"

	"github.com/gin-gonic/gin"
)

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

func NewHandler(c *Config) {
	// Create an group
	g := c.R.Group(os.Getenv("WX_API_URL"))

	uh := &UserHandler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}
	g.GET("/me", uh.Me)
	g.POST("/signup", uh.Signup)
	g.POST("/signin", uh.Signin)
	g.POST("/signout", uh.Signout)
	g.POST("/tokens", uh.Tokens)
	g.POST("/image", uh.Image)
	g.DELETE("/image", uh.DeleteImage)
	g.PUT("/details", uh.Details)
}
