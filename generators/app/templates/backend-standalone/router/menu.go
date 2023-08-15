package router

import (
	v1 "<%= displayName %>/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	MenuRouter := Router.Group("menu")
	{
		MenuRouter.POST("/getMenuByUser", v1.GetMenu)
		MenuRouter.GET("/getMenuList", v1.GetMenuList)
		MenuRouter.POST("/addMenu", v1.AddMenu)
		MenuRouter.DELETE("/deleteMenu", v1.DeleteMenu)
		MenuRouter.POST("/updateMenu", v1.UpdateMenu)
		MenuRouter.GET("/getMenuById", v1.GetMenuById)
		MenuRouter.GET("/getMenuByRole", v1.GetRoleMenu)
		MenuRouter.POST("/addRoleMenu", v1.UpdateRoleMenu)
		MenuRouter.POST("/addButton", v1.AddButton)
		MenuRouter.GET("/getButtonByUser", v1.GetButtonByUser)
		MenuRouter.POST("/updateButton", v1.UpdateButton)
		MenuRouter.POST("/deleteButton", v1.DeleteButton)
	}
}
