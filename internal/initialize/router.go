package initialize

import (
	"manshon/internal/router"

	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Router() *gin.Engine {
	var Router = gin.Default()

	//跨域中间件
	//TODO

	//路由组实例
	serverRouter := router.RouterGroupApp.Server
	PrivateGroup := Router.Group("api/v1.0")
	{
		serverRouter.InitApiRouter(PrivateGroup)
	}

	return Router
}
