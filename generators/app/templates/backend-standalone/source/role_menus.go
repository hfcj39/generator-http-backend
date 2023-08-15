package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var RolesMenus = new(roleMenus)

type roleMenus struct{}

type RoleMenus struct {
	RoleID         uint `gorm:"column:role_id"`
	SysBaseMenusID uint `gorm:"column:sys_base_menu_id"`
}

func genRoleMenus() []RoleMenus {
	var menus []model.SysBaseMenu
	var rolesMenus []RoleMenus
	global.DB.Model(model.SysBaseMenu{}).Find(&menus)
	/// 777 superAdmin
	for _, m := range menus {
		rolesMenus = append(rolesMenus, RoleMenus{1, m.ID})
	}
	return rolesMenus
}

func (r *roleMenus) Init() error {
	rolesMenus := genRoleMenus()
	return global.DB.Table("role_menu").Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rolesMenus).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
			global.LOG.Info("初始化roles_menus表成功")
			return nil
		})
}
