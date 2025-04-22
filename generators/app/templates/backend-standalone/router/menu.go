package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("menu")
	{
		MenuRouter.POST("/list", v1.GetMenuList) // 管理页面用，获取全部的路由
		MenuRouter.GET("/all", v1.GetMenu)       // 根据用户token获取对应的页面路由
		MenuRouter.POST("/addMenu", v1.AddMenu)
		MenuRouter.DELETE("/deleteMenu", v1.DeleteMenu)
		MenuRouter.POST("/updateMenu", v1.UpdateMenu)
		MenuRouter.GET("/getMenuById", v1.GetMenuById)
		MenuRouter.GET("/getMenuByRole", v1.GetRoleMenu)
		MenuRouter.POST("/addRoleMenu", v1.UpdateRoleMenu)
	}
}
