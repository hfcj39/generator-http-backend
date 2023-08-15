package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var layout = "/@/layouts/default/index.vue"

var Menus = []model.SysBaseMenu{
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "dashboard",
		Name:       "Dashboard",
		Redirect:   "/dashboard/analysis",
		Component:  layout,
		Meta:       model.Meta{Title: "routes.dashboard.dashboard", Icon: "ion:grid-outline", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "Dashboard",
		Path:       "analysis",
		Name:       "DashboardAnalysis",
		Component:  "/@/views/dashboard/analysis/index.vue",
		Meta:       model.Meta{Title: "routes.dashboard.analysis", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "user",
		Name:       "User",
		Redirect:   "/user/me",
		Component:  layout,
		Meta:       model.Meta{Title: "routes.user.userManagement", Icon: "ant-design:user-outlined", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "User",
		Path:       "me",
		Name:       "UserMe",
		Component:  "/@/views/user/me/index.vue",
		Meta:       model.Meta{Title: "routes.user.me", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "User",
		Path:       "userList",
		Name:       "UserList",
		Component:  "/@/views/user/list/index.vue",
		Meta:       model.Meta{Title: "routes.user.userList", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "system",
		Name:       "System",
		Component:  layout,
		Meta:       model.Meta{Title: "routes.system.systemManagement", Icon: "ion:settings-outline", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "System",
		Path:       "systemSettings",
		Name:       "SystemInfo",
		Component:  "/@/views/system/info/index.vue",
		Meta:       model.Meta{Title: "routes.system.systemSettings", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "User",
		Path:       "role",
		Name:       "UserRole",
		Component:  "/@/views/user/role/index.vue",
		Meta:       model.Meta{Title: "routes.user.role", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "System",
		Path:       "systemMenus",
		Name:       "SystemMenu",
		Component:  "/@/views/system/menu/index.vue",
		Meta:       model.Meta{Title: "routes.system.systemMenus", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "System",
		Path:       "casbin",
		Name:       "SystemCasbin",
		Component:  "/@/views/system/casbin/index.vue",
		Meta:       model.Meta{Title: "routes.system.casbin", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "setup",
		Name:       "Setup",
		Component:  layout,
		Meta:       model.Meta{Title: "routes.setup.setup", IgnoreKeepAlive: &[]bool{true}[0]},
	},
	{
		OrderNo:    0,
		ParentName: "Setup",
		Path:       "demo",
		Name:       "DemoSetup",
		Component:  "/@/views/demo/setup/index.vue",
		Meta:       model.Meta{Title: "routes.setup.setup", IgnoreKeepAlive: &[]bool{true}[0]},
	},
}

func (m *menu) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&Menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		global.LOG.Info("初始化base_menu表成功")
		return nil
	})
}
