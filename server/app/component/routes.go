package component

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	Router  *gin.Engine
	Request *http.Request
)

func GetCurrentRoute() map[string]string {
	var route map[string]string
	if Request == nil {
		return route
	}
	return map[string]string{
		"url":    Request.RequestURI,
		"method": Request.Method,
	}
}

func GetAllRoutes() []map[string]string {
	routers := []map[string]string{}
	if Router == nil {
		return routers
	}
	appRouters := Router.Routes()
	for _, route := range appRouters {
		name := strings.Replace(route.Path, os.Getenv("APP_API_V1"), "", 1)
		tmpList := strings.Split(name, "/")
		name = tmpList[0]
		routers = append(routers, map[string]string{
			"url":    route.Path,
			"name":   name,
			"method": route.Method,
		})
	}
	return routers
}
