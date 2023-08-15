package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var buttons = []model.Button{
	{
		SysBaseMenuName: "UserMe",
		AuthFlag:        "UserMe:edit_password",
	},
	{
		SysBaseMenuName: "UserList",
		AuthFlag:        "UserList:edit",
	},
	{
		SysBaseMenuName: "UserRole",
		AuthFlag:        "UserRole:edit",
	},
	{
		SysBaseMenuName: "UserRole",
		AuthFlag:        "UserRole:delete",
	},
	{
		SysBaseMenuName: "UserRole",
		AuthFlag:        "UserRole:create",
	},
	{
		SysBaseMenuName: "UserRole",
		AuthFlag:        "UserRole:menu",
	},
	{
		SysBaseMenuName: "SystemCasbin",
		AuthFlag:        "SystemCasbin:auth",
	},
	{
		SysBaseMenuName: "SystemCasbin",
		AuthFlag:        "SystemCasbin:method",
	},
	{
		SysBaseMenuName: "SystemMenu",
		AuthFlag:        "SystemMenu:edit",
	},
	{
		SysBaseMenuName: "SystemMenu",
		AuthFlag:        "SystemMenu:create",
	},
	{
		SysBaseMenuName: "SystemMenu",
		AuthFlag:        "SystemMenu:delete",
	},
	{
		SysBaseMenuName: "SystemMenu",
		AuthFlag:        "SystemMenu:create_child",
	},
	{
		SysBaseMenuName: "SystemCasbin",
		AuthFlag:        "SystemCasbin:button",
	},
	{
		SysBaseMenuName: "ServerConfig",
		AuthFlag:        "ServerConfig:cascade_edit",
	},
	{
		SysBaseMenuName: "ServerConfig",
		AuthFlag:        "ServerConfig:edit_description",
	},
	{
		SysBaseMenuName: "ServerConfig",
		AuthFlag:        "ServerConfig:add",
	},
	{
		SysBaseMenuName: "ServerConfig",
		AuthFlag:        "ServerConfig:remove",
	},
}

type button struct{}

var Button = new(button)

func (a *button) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&buttons).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("初始化button表成功")
		return nil
	})
}
