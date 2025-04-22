package model

// 取消，button直接合并到menu中

// type Button struct {
// 	global.BASE_MODEL
// 	SysBaseMenuName string `json:"sys_base_menu_name" gorm:"not null;default:0"`
// 	AuthFlag        string `json:"auth_flag" gorm:"comment:权限标识;unique"`
// 	IsActive        *bool  `json:"is_active" gorm:"comment:是否启用;default:true;not null"`
// 	Type            string `json:"type" gorm:"comment:按钮权限类型(casbin/menu/token);not null;default:casbin"`
// 	ToMenuName      string `json:"to_menu_name" gorm:"comment:跳转menu按钮的menu_name"`
// 	IsSameAuthority *bool  `json:"is_same_authority" gorm:"comment:是否需要同一auth才可操作;default:false;not null"`
// 	OrderNo         int    `json:"order_no" gorm:"default:0;comment:排序标记"`
// 	Description     string `json:"description" gorm:"comment:描述"`
// 	RoleValue       int    `json:"role_value" gorm:"comment:其他接口所需的角色值"`
// }
