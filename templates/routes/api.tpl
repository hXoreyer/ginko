package routes

import (
	"{{.Name}}/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func SetV1GroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pongv1")
	})
	router.POST("/register", user.Register)
}
