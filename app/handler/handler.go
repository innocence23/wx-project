package handler

import (
	"net/http"
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

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":   0,
		"status": "success",
		"smg":    "成功",
		"data":   data,
	})
}

func Fail(ctx *gin.Context, errcode int, errmsg ...interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":   errcode,
		"status": "fail",
		"msg":    errmsg,
		"data":   nil,
	})
}
