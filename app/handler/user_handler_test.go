package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Me(t *testing.T) {
	//创建一个请求
	req, err := http.NewRequest(http.MethodGet, "/user/info", nil)
	assert.NoError(t, err)

	//我们创建一个 ResponseRecorder 来记录响应
	rr := httptest.NewRecorder()
	router := gin.Default()
	NewHandler(&Config{
		R: router,
	})
	router.ServeHTTP(rr, req)

	// 检测返回的状态码
	assert.Equal(t, 200, rr.Code)
	// 检测返回的数据
	var respBody = "{\"code\":400,\"data\":null,\"msg\":{\"error\":{\"type\":\"BADREQUEST\",\"message\":\"请求头中auth为空\"}},\"status\":\"fail\"}"
	assert.Equal(t, respBody, rr.Body.String())
}
