package request

type AddRoleStruct struct {
	RoleName          string  `json:"role_name" binding:"required"`
	RoleValue         int     `json:"role_value" binding:"required"`
	Description       string  `json:"description"`
	ButtonPermissions []int64 `json:"button_permissions"`
}

type UpdateRoleStruct struct {
	ID                uint    `json:"id" binding:"required"`
	RoleName          string  `json:"role_name" binding:"required"`
	RoleValue         int     `json:"role_value" binding:"required"`
	Description       string  `json:"description"`
	ButtonPermissions []int64 `json:"button_permissions"`
}

type GetRoleListStruct struct {
	RoleName   string `form:"role_name"`
	FilterMode string `form:"filter_mode"`
	ListParamsStruct
}

type GetRoleUserListStruct struct {
	RoleID uint `form:"role_id" binding:"required"`
	ListParamsStruct
}

type AddRoleUserStruct struct {
	RoleID  uint   `json:"role_id" binding:"required"`
	UserIDs []uint `json:"user_ids" binding:"required"`
}
