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

var dashboardMenus = []model.SysBaseMenu{
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "dashboard",
		Name:       "Dashboard",
		Redirect:   "/dashboard/analysis",
		Component:  layout,
		Type:       model.MenuTypeCatalog,
		Meta:       model.Meta{Title: "routes.dashboard.dashboard", Icon: "ion:grid-outline"},
	},
	{
		OrderNo:    0,
		ParentName: "Dashboard",
		Path:       "analysis",
		Name:       "DashboardAnalysis",
		Component:  "/@/views/dashboard/analysis/index.vue",
		Type:       model.MenuTypeMenu,
		Meta:       model.Meta{Title: "routes.dashboard.analysis"},
	},
}

var userMenus = []model.SysBaseMenu{
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "user",
		Name:       "User",
		Redirect:   "/user/me",
		Component:  layout,
		Type:       model.MenuTypeCatalog,
		Meta:       model.Meta{Title: "routes.user.userManagement", Icon: "ant-design:user-outlined"},
	},
	{
		OrderNo:    0,
		ParentName: "User",
		Path:       "me",
		Name:       "UserMe",
		Component:  "/@/views/user/me/index.vue",
		Type:       model.MenuTypeMenu,
		Meta:       model.Meta{Title: "routes.user.me"},
	},
	{
		OrderNo:    0,
		ParentName: "User",
		Path:       "userList",
		Name:       "UserList",
		Component:  "/@/views/user/list/index.vue",
		Type:       model.MenuTypeMenu,
		Meta:       model.Meta{Title: "routes.user.userList"},
	},
}

var systemMenus = []model.SysBaseMenu{
	{
		OrderNo:    0,
		ParentName: "0",
		Path:       "system",
		Name:       "System",
		Component:  layout,
		Type:       model.MenuTypeCatalog,
		Meta:       model.Meta{Title: "routes.system.systemManagement", Icon: "ion:settings-outline"},
	},
	{
		OrderNo:    0,
		ParentName: "System",
		Path:       "systemSettings",
		Name:       "SystemInfo",
		Component:  "/@/views/system/info/index.vue",
		Type:       model.MenuTypeMenu,
		Meta:       model.Meta{Title: "routes.system.systemSettings"},
	},
	{
		OrderNo:    0,
		ParentName: "System",
		Path:       "systemMenus",
		Name:       "SystemMenu",
		Component:  "/@/views/system/menu/index.vue",
		Type:       model.MenuTypeMenu,
		Meta:       model.Meta{Title: "routes.system.systemMenus"},
	},
	{
		OrderNo:    0,
		ParentName: "System",
		Path:       "casbin",
		Name:       "SystemCasbin",
		Component:  "/@/views/system/casbin/index.vue",
		Type:       model.MenuTypeMenu,
		Meta:       model.Meta{Title: "routes.system.casbin"},
	},
}

// Init方法
func (m *menu) Init() error {
	// 将所有菜单数组存放在一个数组里
	var menuArrays = [][]model.SysBaseMenu{
		dashboardMenus,
		userMenus,
		systemMenus,
	}

	// 使用一次 append 将所有菜单合并成一个数组
	var Menus []model.SysBaseMenu
	for _, menuArray := range menuArrays {
		Menus = append(Menus, menuArray...)
	}

	// 在事务中插入数据
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
