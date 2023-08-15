package initialize

import (
	_ "<%= displayName %>/docs"
	"<%= displayName %>/global"
	"<%= displayName %>/middleware"
	"<%= displayName %>/router"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.CONFIG.System.StaticPath, http.Dir(global.CONFIG.System.StaticPath)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.CORS()).Use(middleware.VerifyJwtToken()).Use(middleware.CasbinHandler())
	{
		router.InitUserRouter(PrivateGroup)
		router.InitMenuRouter(PrivateGroup)
		router.InitRoleRouter(PrivateGroup)
		router.InitCasbinRouter(PrivateGroup)
		router.InitSystemRouter(PrivateGroup)
	}
	global.LOG.Info("router register success")
	return Router
}
