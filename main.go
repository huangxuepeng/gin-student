package main

import (
	"student/initialize"

	"go.uber.org/zap"
)

func main() {
	port := "8080"
	//初始化logger
	initialize.InitLogger()
	//初始化router
	router := initialize.Routers()
	zap.S().Debugf("启动服务器, 端口: " + port)
	if err := router.Run(":8081"); err != nil {
		zap.S().Panic("启动失败: ", err.Error())
	}
}
