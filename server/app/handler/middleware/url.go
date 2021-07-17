package middleware

import (
	"wx/app/component"

	"github.com/gin-gonic/gin"
)

func UrlPathMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		component.Request = ctx.Request
		ctx.Next()
	}
}
