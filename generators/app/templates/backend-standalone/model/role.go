// buttonPermissions []

package model

import (
	"<%= displayName %>/global"

	pg "github.com/lib/pq"
)

type Role struct {
	global.BASE_MODEL
	RoleName          string        `json:"role_name" gorm:"comment:角色名称;unique;not null"`
	RoleValue         int           `json:"role_value" gorm:"comment:角色值;unique;not null"`
	Description       string        `json:"description" gorm:"comment:描述"`
	Sort              int           `json:"sort" gorm:"comment:排序"`
	ButtonPermissions pg.Int64Array `json:"-" gorm:"type:int[];comment:按钮权限值"` // todo
	SysBaseMenus      []SysBaseMenu `json:"sys_base_menus,omitempty" gorm:"many2many:role_menu;"`
}
