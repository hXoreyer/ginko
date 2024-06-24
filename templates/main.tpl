package main

import (
	"{{.Name}}/bootstrap"
	"{{.Name}}/global"
)

// @title Gin Swagger Example API
// @description This is a sample server for a Gin application.
// @version 1.0

// @host localhost:8888

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @description Type into the textbox: Bearer {your jwt}.


func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 启动服务器
	bootstrap.RunServer()
}
