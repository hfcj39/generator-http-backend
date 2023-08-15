package model

import (
	"<%= displayName %>/global"
)

type SysBaseMenu struct {
	global.BASE_MODEL
	ParentName string `json:"parent_name" binding:"required" gorm:"default:0;comment:父菜单名称"`
	Path       string `json:"path" binding:"required" gorm:"comment:路由path;unique"`
	Name       string `json:"name" binding:"required" gorm:"comment:路由name;unique"`
	Component  string `json:"component" binding:"required" gorm:"comment:对应前端文件路径;not null;"`
	Redirect   string `json:"redirect" gorm:"comment:跳转路径"`
	OrderNo    int    `json:"order_no" gorm:"default:0;comment:排序标记"`
	Meta       `json:"meta" gorm:"comment:附加属性"`
	// Role       []Role                 `gorm:"many2many:role_menus;"`
	Children   []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters []SysBaseMenuParameter `json:"parameters"`
	Buttons    []Button               `json:"buttons" gorm:"foreignKey:SysBaseMenuName;references:Name"`
}

type Meta struct {
	IgnoreKeepAlive    *bool  `json:"ignoreKeepAlive" gorm:"column:ignoreKeepAlive;comment:是否缓存"`
	Title              string `json:"title" binding:"required" gorm:"comment:菜单名"`
	Icon               string `json:"icon" gorm:"comment:菜单图标"`
	HideMenu           *bool  `json:"hideMenu" gorm:"column:hideMenu;default:false;comment:Never show in menu"`
	HideTab            *bool  `json:"hideTab" gorm:"column:hideTab;comment:是否隐藏tab"`
	TransitionName     string `json:"transitionName" gorm:"column:transitionName;comment:current page transition"`
	Affix              *bool  `json:"affix" gorm:"comment:Is it fixed on tab"`
	FrameSrc           string `json:"frameSrc" gorm:"column:frameSrc;"`
	HideBreadcrumb     *bool  `json:"hideBreadcrumb" gorm:"comment:Whether the route has been dynamically added"`
	HideChildrenInMenu *bool  `json:"hideChildrenInMenu" gorm:"comment:Hide submenu"`
	CarryParam         *bool  `json:"carryParam" gorm:"comment:Carrying parameters"`
	Single             *bool  `json:"single" gorm:"comment:Used internally to mark single-level menus"`
	IsLink             *bool  `json:"isLink"`
	HiddenFooter       *bool  `json:"hiddenFooter" gorm:"column:hiddenFooter;comment:是否隐藏footer"`
	Description        string `json:"description" gorm:"column:description;comment:菜单描述"`
}

type SysBaseMenuParameter struct {
	global.BASE_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}
