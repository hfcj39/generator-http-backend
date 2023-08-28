package request

// OidcCallback structure
type OidcCallback struct {
	Code  string `json:"code" binding:"required"`
	State string `json:"state"`
}

type UpdateUserStruct struct {
	ID          uint      `json:"id" binding:"required"`
	RoleID      uint      `json:"role_id"`
	CustomGroup string    `json:"custom_group"`
	Remark      string    `json:"remark"`
	HeaderImg   string    `json:"header_img"`
	DisplayName string    `json:"display_name"`
}

type UpdateSelfUserInfoStruct struct {
	Remark      string `json:"remark" binding:"required"`
	HeaderImg   string `json:"header_img" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
}

type GetUserListStruct struct {
	Name           string `form:"name"`
	Role           string `form:"role"`
	RoleID         uint   `form:"role_id"`
	ExcludedRoleID uint   `form:"excluded_role_id"`
	ListParamsStruct
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SetUserPassword struct {
	OldPassword    *string `json:"old_password"` // 增加密码不传该字段, 修改密码填写该字段
	Password       string  `json:"password" binding:"required"`
	RepeatPassword string  `json:"repeat_password" binding:"required,eqfield=Password"`
}

type GetLDAPUserListStruct struct {
	Name string `form:"name" binding:"required,min=2"` // 搜索人名必填,且大于两个个字母才能搜索
}
