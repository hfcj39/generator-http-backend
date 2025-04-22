package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"
	"<%= displayName %>/utils"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

//TODO 初始化admin帐号

var Admin = new(admin)

type admin struct{}

var admins = []model.User{
	{
		BASE_MODEL: global.BASE_MODEL{ID: 1},
		Username:   "sa",
		Password:   utils.MD5V("hello123"),
		RealName:   "sa",
		RoleID:     1,
	},
	{
		BASE_MODEL: global.BASE_MODEL{ID: 2},
		Username:   "admin",
		Password:   utils.MD5V("hello123"),
		RealName:   "admin",
		RoleID:     2,
	},
}

func (a *admin) Init() error {
	return global.DB.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Clauses(clause.OnConflict{
				DoNothing: true,
			}).Create(&admins).Error; err != nil { // 遇到错误时回滚事务
				return err
			}
			if err := tx.Exec("SELECT setval('users_id_seq',MAX(id),true) FROM users").Error; err != nil {
				return err
			}
			global.LOG.Info("初始化user表成功")
			return nil
		})
}
