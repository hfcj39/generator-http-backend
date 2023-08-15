package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("role")
	{
		RoleRouter.GET("/list", v1.GetRoleList)
		RoleRouter.GET("/detail/:id", v1.FindRoleById)
		RoleRouter.POST("/create", v1.CreateRole)
		RoleRouter.POST("/update", v1.UpdateRole)
		RoleRouter.DELETE("/:id", v1.DeleteRole)
	}

	RoleUserRouter := Router.Group("roleuser")

	RoleUserRouter.GET("/list", v1.GetRoleUserList)
	RoleUserRouter.POST("/add", v1.AddRoleUser)
	RoleUserRouter.DELETE("/:id", v1.DeleteRoleUser)
}
