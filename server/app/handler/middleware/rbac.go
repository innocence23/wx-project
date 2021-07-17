package middleware

import (
	"wx/app/component"
	"wx/app/dto"
	"wx/app/handler/commonhandler"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
)

func RbacMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			msg := "上下文中获取不到用户"
			e := zerror.NewAuthorization(msg)
			commonhandler.Fail(c, e.Status(), gin.H{
				"error": e,
			})
			c.Abort()
		}
		name := user.(*dto.UserJWT).Email
		permission := c.FullPath()
		method := c.Request.Method
		if !component.CheckPermission(name, permission, method) {
			msg := "您没有权限操作"
			e := zerror.NewAuthorization(msg)
			commonhandler.Fail(c, e.Status(), gin.H{
				"error": e,
			})
			c.Abort()
		}
		c.Next()
	}
}
