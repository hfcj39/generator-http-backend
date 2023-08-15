package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Roles = []model.Role{
	{BASE_MODEL: global.BASE_MODEL{ID: 1}, RoleName: "superAdmin", RoleValue: 777, Description: "超级管理员", Sort: 0, ButtonPermissions: []int64{}},
	{BASE_MODEL: global.BASE_MODEL{ID: 2}, RoleName: "admin", RoleValue: 555, Description: "管理员", Sort: 0, ButtonPermissions: []int64{}},
	{BASE_MODEL: global.BASE_MODEL{ID: 3}, RoleName: "validUser", RoleValue: 444, Description: "用户", Sort: 0, ButtonPermissions: []int64{}},
}

type role struct{}

var Role = new(role)

func (r *role) Init() error {
	return global.DB.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&Roles).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
			if err := tx.Exec("SELECT setval('roles_id_seq',MAX(id),true) FROM roles").Error; err != nil {
				return err
			}
			global.LOG.Info("初始化roles表成功")
			return nil
		})
}
