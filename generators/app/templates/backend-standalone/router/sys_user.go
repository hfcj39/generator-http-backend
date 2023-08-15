package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("/userInfo", v1.GetUserInfo)
		UserRouter.DELETE("/:id", v1.DeleteUser)
		UserRouter.GET("/list", v1.GetUserList)
		UserRouter.POST("/update", v1.UpdateUser)
		UserRouter.GET("/detail/:id", v1.FindUserById)
		UserRouter.POST("/password", v1.SetPassword)
		UserRouter.POST("/userInfo/update", v1.UpdateSelfUserInfo)
	}
}
