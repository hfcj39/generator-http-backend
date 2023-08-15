package main

import (
	"<%= displayName %>/core"
	"<%= displayName %>/global"
	"<%= displayName %>/initialize"
)

// @title <%= displayName %>
// @version 0.0
// @description 冲冲冲
// @contact.name hfcj
// @BasePath /
func main() {
	global.VP = core.Viper()      // 初始化Viper
	core.Env()                    // 通过环境变量配置参数
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	if global.DB != nil {
		initialize.PostgresTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	initialize.Redis() // redis

	core.RunServer()
}
