package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		BaseRouter.POST("/user/login", v1.Login)
		BaseRouter.GET("/version", v1.GetSystemVersion)
	}
	return BaseRouter
}
