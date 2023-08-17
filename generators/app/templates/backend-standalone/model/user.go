package model

import (
	"<%= displayName %>/global"

	"github.com/volatiletech/null"

	"github.com/lib/pq"
)

type User struct {
	global.BASE_MODEL
	Username    string         `json:"username" gorm:"comment:用户登录名;unique;not null"`
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`
	DisplayName string         `json:"display_name"  gorm:"comment:ldap中displayName字段"`
	CustomGroup string         `json:"custom_group"  gorm:"comment:自定义组"`
	Remark      string         `json:"remark"  gorm:"comment:备注"`
	HeaderImg   string         `json:"header_img" gorm:"comment:用户头像"`
	RoleID      uint           `json:"role_id" gorm:"comment:用户角色ID"`
	Role        *Role          `json:"role,omitempty"`

	AccessToken  string `json:"-" gorm:"comment:access_token"`
	RefreshToken string `json:"-" gorm:"comment:refresh_token"`
	IdToken      string `json:"-" gorm:"comment:id_token"`
}
