// buttonPermissions []

package model

import (
	"<%= displayName %>/global"
)

type Role struct {
	global.BASE_MODEL
	RoleName          string        `json:"roleName" gorm:"comment:角色名称;unique;not null"`
	RoleValue         int           `json:"roleValue" gorm:"comment:角色值;unique;not null"` // 决定了能访问哪些后段接口
	Description       string        `json:"description" gorm:"comment:描述"`
	Sort              int           `json:"sort" gorm:"comment:排序"`
	ButtonPermissions []uint        `json:"-" gorm:"-"` // 取消，permissions字段直接返回有权限的menuid
	SysBaseMenus      []SysBaseMenu `json:"sysBaseMenus,omitempty" gorm:"many2many:role_menu;"`
}
