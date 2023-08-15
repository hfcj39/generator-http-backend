package middleware

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model/response"
	"<%= displayName %>/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(RoleClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := fmt.Sprint(waitUse.RoleValue)
		e := service.Casbin()

		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.CONFIG.System.Env == "test" || success {
			c.Next()
		} else {
			response.FailWithMessage("Casbin 权限不足", c)
			c.Abort()
			return
		}
	}
}
