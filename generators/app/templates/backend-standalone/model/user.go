package model

import (
	"<%= displayName %>/global"
)

type User struct {
	global.BASE_MODEL
	Username    string `json:"username" gorm:"comment:用户登录名;unique;not null"`
	Password    string `json:"-" gorm:"comment:用户登录密码"`
	RealName    string `json:"realName" gorm:"comment:displayName"`
	CustomGroup string `json:"customGroup" gorm:"comment:自定义组"`
	HomePath    string `json:"homePath" gorm:"comment:首页路径"`
	Remark      string `json:"remark" gorm:"comment:备注"`
	Description string `json:"description" gorm:"comment:描述"`
	Avatar      string `json:"avatar" gorm:"comment:用户头像"`
	RoleID      uint   `json:"roleId" gorm:"comment:用户角色ID"`
	Role        *Role  `json:"role,omitempty"`
}
