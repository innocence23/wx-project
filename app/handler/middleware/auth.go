package middleware

import (
	"log"
	"strings"
	"wx/app/component"
	"wx/app/handler/commonhandler"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			msg := "请求头中auth为空"
			e := zerror.NewBadRequest(msg)
			commonhandler.Fail(c, e.Status(), gin.H{
				"error": e,
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			msg := "请求头中auth格式有误"
			e := zerror.NewBadRequest(msg)
			commonhandler.Fail(c, e.Status(), gin.H{
				"error": e,
			})
			c.Abort()
			return
		}
		result, err := component.ParseToken(parts[1])
		if err != nil {
			log.Printf("无效的Token: %#v\n", err)
			msg := "无效的Token"
			e := zerror.NewBadRequest(msg)
			commonhandler.Fail(c, e.Status(), gin.H{
				"error": e,
			})
			c.Abort()
			return
		}
		c.Set("user", result.User)
		c.Next()
	}
}
