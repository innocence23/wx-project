package handler

import (
	"net/http"
	"os"
	"wx/app/model"

	"github.com/gin-gonic/gin"
)

type Config struct {
	R           *gin.Engine
	UserService model.UserService
}

func NewHandler(c *Config) {
	g := c.R.Group(os.Getenv("WX_API_URL_V1"))

	uh := &UserHandler{
		UserService: c.UserService,
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

func success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":   0,
		"status": "success",
		"smg":    "成功",
		"data":   data,
	})
}

func fail(ctx *gin.Context, errcode int, errmsg ...interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":   errcode,
		"status": "fail",
		"msg":    errmsg,
		"data":   nil,
	})
}
