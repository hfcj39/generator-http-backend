package middleware

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model/response"
	"<%= displayName %>/service"
	"<%= displayName %>/utils"
	"<%= displayName %>/utils/e"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RoleClaims struct {
	RoleValue int
	utils.Claims
}

func VerifyJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) <= 0 {
			response.FailWithDetailed(e.ErrorAuth, "获取token失败", e.GetMsg(e.ErrorAuth), c)
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		claims, err := utils.ParseToken(token)
		if err != nil {
			var code int
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = e.ErrorAuthCheckTokenTimeout
			default:
				code = e.ErrorAuthCheckTokenFail
			}
			response.FailWithDetailed(code, "", e.GetMsg(code), c)
			c.Abort()
			return
		}
		// err, r := service.FindRoleById(claims.RoleID)
		err, v := service.GetRoleValueByUserId(claims.UserId)
		if err != nil {
			global.LOG.Error(err.Error())
			response.FailWithMessage("获取角色权限值失败,请检查登录用户", c)
			c.Abort()
			return
		}
		c.Set("claims", RoleClaims{v, *claims})
		c.Next()
	}
}
