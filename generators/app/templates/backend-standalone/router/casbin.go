package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	GroupRouter := Router.Group("casbin")
	{
		GroupRouter.GET("/list", v1.GetCasbinList)
		GroupRouter.POST("/update", v1.UpdateCasbinRule)
		GroupRouter.POST("/updateById", v1.UpdateCasbinRuleById)
	}
}
