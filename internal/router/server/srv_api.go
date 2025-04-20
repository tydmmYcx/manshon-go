package server

import (
	"manshon-go/internal/api"

	"github.com/gin-gonic/gin"
)

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	var apiRouterApi = api.ApiGroupApp.ServerApiGroup.ServerApiGroup
	{
		Router.POST("/", apiRouterApi.ReflectAction)
		Router.POST("/report", apiRouterApi.Report)
	}
}
