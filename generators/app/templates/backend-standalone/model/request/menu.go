package request

type AddRoleMenu struct {
	Menus  []uint `json:"menus" binding:"required"`
	RoleId int    `json:"role_id" binding:"required"`
}

type AddButton struct {
	SysBaseMenuName string `json:"sys_base_menu_name" binding:"required"`
	AuthFlag        string `json:"auth_flag" binding:"required"`
	IsActive        *bool  `json:"is_active" binding:"required"`
	OrderNo         *int   `json:"order_no" binding:"required"`
	Type            string `json:"type" binding:"required"`
	IsSameAuthority *bool  `json:"is_same_authority" binding:"required"`
	Description     string `json:"description" binding:"required"`
	RoleValue       int    `json:"role_value" binding:"required"`
}

type UpdateButton struct {
	ID              uint    `json:"id" binding:"required"`
	SysBaseMenuName string  `json:"sys_base_menu_name" binding:"required"`
	AuthFlag        string  `json:"auth_flag" binding:"required"`
	IsActive        *bool   `json:"is_active" binding:"required"`
	OrderNo         *int    `json:"order_no" binding:"required"`
	Type            string  `json:"type" binding:"required"`
	IsSameAuthority *bool   `json:"is_same_authority" binding:"required"`
	Description     *string `json:"description" binding:"required"`
}
