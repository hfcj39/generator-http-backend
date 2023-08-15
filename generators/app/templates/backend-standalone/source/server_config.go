package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"

	"gorm.io/datatypes"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

var configs = []model.ServerConfig{
	{
		BASE_MODEL:  global.BASE_MODEL{ID: 2},
		ConfigName:  "server",
		Description: "服务器名",
		ConfigType:  "string",
		ConfigValue: datatypes.JSON(`{"value": "server"}`),
	},
}

// todo 修改版本号三种方式
// 1. 发版本之前修改source以修改数据库的值
// 2. 通过update config接口修改,但是OnConflict时会被旧值覆盖
// 3. 执行sql
var versionConfigs = []model.ServerConfig{
	{
		BASE_MODEL:  global.BASE_MODEL{ID: 1},
		ConfigName:  "systemVersion",
		Description: "本系统版本号",
		ConfigType:  "string",
		ConfigValue: datatypes.JSON(`{"value": "0.0.0"}`),
		ConfigScope: "system",
	},
}

type config struct{}

var Config = new(config)

func (c *config) Init() error {
	return global.DB.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(&configs).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
			if err := tx.Clauses(clause.OnConflict{
				UpdateAll: true, // 版本号需一直更新
			}).Create(&versionConfigs).Error; err != nil {
				return err
			}
			if err := tx.Exec("SELECT setval('server_configs_id_seq',MAX(id),true) FROM server_configs").Error; err != nil {
				return err
			}
			global.LOG.Info("初始化server_configs表成功")
			return nil
		})
}
