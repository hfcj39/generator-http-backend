package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system")
	{
		SystemRouter.GET("/getSystemConfig", v1.GetSystemConfig)
	}
}
