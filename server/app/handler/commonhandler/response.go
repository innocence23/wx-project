package commonhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"status":  "success",
		"message": "成功",
		"data":    data,
	})
}

func Fail(ctx *gin.Context, errcode int, errmsg interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    errcode,
		"status":  "fail",
		"message": errmsg,
		"data":    nil,
	})
}
