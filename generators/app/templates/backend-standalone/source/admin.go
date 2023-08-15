package source

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"
	"<%= displayName %>/utils"

	"github.com/volatiletech/null"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

//TODO 初始化admin帐号

var Admin = new(admin)

type admin struct{}

var admins = []model.User{
	{
		BASE_MODEL:  global.BASE_MODEL{ID: 1},
		Username:    "huangfuchenjie",
		Password:    "e10adc3949ba59abbe56e057f20f883e",
		DisplayName: "hfcj",
		AuthorityID: null.Uint{
			Uint: 1, Valid: true,
		},
		RoleID: 1,
	},
	{
		Username:    "admin",
		Password:    utils.MD5V("hellocnops12345"),
		DisplayName: "admin",
		AuthorityID: null.Uint{
			Uint: 1, Valid: true,
		},
		RoleID: 2,
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
