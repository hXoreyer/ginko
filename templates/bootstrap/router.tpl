package bootstrap

import (
	"{{.Name}}/global"
	"{{.Name}}/middlewares"
	"{{.Name}}/routes"

	"github.com/gin-gonic/gin"

  	_ "{{.Name}}/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middlewares.Cors())
	groups := routes.Groups(router)
	groups.Init()
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.App.Port)
}
