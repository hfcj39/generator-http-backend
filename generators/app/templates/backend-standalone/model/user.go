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
	LdapUID     string         `json:"ldap_uid"  gorm:"comment:ldap中uid字段"`
	LdapGroup   pq.StringArray `json:"ldap_group"  gorm:"comment:ldap中group字段;type:varchar[]"`
	LdapMail    string         `json:"ldap_mail"  gorm:"comment:ldap中mail字段"`
	CustomGroup string         `json:"custom_group"  gorm:"comment:自定义组"`
	Remark      string         `json:"remark"  gorm:"comment:备注"`
	HeaderImg   string         `json:"header_img" gorm:"comment:用户头像"`
	AuthorityID null.Uint      `json:"authority_id" gorm:"comment:用户管理服务器权限ID"`
	RoleID      uint           `json:"role_id" gorm:"comment:用户角色ID"`
	Role        *Role          `json:"role,omitempty"`

	AccessToken  string `json:"-" gorm:"comment:access_token"`
	RefreshToken string `json:"-" gorm:"comment:refresh_token"`
	IdToken      string `json:"-" gorm:"comment:id_token"`
}
