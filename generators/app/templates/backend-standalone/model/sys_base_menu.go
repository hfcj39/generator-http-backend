package model

import (
	"<%= displayName %>/global"
	"database/sql/driver"
	"fmt"

	"gorm.io/datatypes"
)

type MenuType string

const (
	MenuTypeCatalog  MenuType = "catalog"
	MenuTypeMenu     MenuType = "menu"
	MenuTypeEmbedded MenuType = "embedded"
	MenuTypeLink     MenuType = "link"
	MenuTypeButton   MenuType = "button"
)

func (t MenuType) IsValid() bool {
	switch t {
	case MenuTypeCatalog, MenuTypeMenu, MenuTypeEmbedded, MenuTypeLink, MenuTypeButton:
		return true
	}
	return false
}

// 保证数据正确用，避免插入其他值
func (t *MenuType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid enum value: %v", value)
	}
	*t = MenuType(str)
	return nil
}

func (t MenuType) Value() (driver.Value, error) {
	if !t.IsValid() {
		return nil, fmt.Errorf("invalid enum value: %v", t)
	}
	return string(t), nil
}

type SysBaseMenu struct {
	global.BASE_MODEL
	Meta       `json:"meta" gorm:"comment:附加属性"`
	ParentName string                 `json:"parentName" binding:"required" gorm:"column:parent_name;default:0;comment:父菜单名称"`
	PID        uint                   `json:"pid" gorm:"column:pid;comment:父级菜单ID"`
	Path       string                 `json:"path" binding:"required" gorm:"column:path;comment:路由path;unique"`
	Name       string                 `json:"name" binding:"required" gorm:"column:name;comment:路由name;unique"`
	Component  string                 `json:"component" binding:"required" gorm:"column:component;comment:对应前端文件路径;not null;"`
	Redirect   string                 `json:"redirect" gorm:"column:redirect;comment:跳转路径"`
	OrderNo    int                    `json:"orderNo" gorm:"column:order_no;default:0;comment:排序标记"`
	Children   []SysBaseMenu          `json:"children" gorm:"-"`
	Parameters []SysBaseMenuParameter `json:"parameters" gorm:"-"`
	Type       MenuType               `json:"type" binding:"required" gorm:"column:type;comment:菜单类型"`
	Status     int                    `json:"status" gorm:"column:status;comment:菜单状态"`
}
type Meta struct {
	ActiveIcon         string         `json:"activeIcon" gorm:"column:active_icon;comment:激活时显示的图标"`
	ActivePath         string         `json:"activePath" gorm:"column:active_path;comment:作为路由时，需要激活的菜单的Path"`
	AffixTab           *bool          `json:"affixTab" gorm:"column:affix_tab;comment:固定在标签栏"`
	AffixTabOrder      *int           `json:"affixTabOrder" gorm:"column:affix_tab_order;comment:在标签栏固定的顺序"`
	Badge              string         `json:"badge" gorm:"column:badge;comment:徽标内容(当徽标类型为normal时有效)"`
	BadgeType          string         `json:"badgeType" gorm:"column:badge_type;comment:徽标类型"`
	BadgeVariants      string         `json:"badgeVariants" gorm:"column:badge_variants;comment:徽标颜色"`
	HideChildrenInMenu *bool          `json:"hideChildrenInMenu" gorm:"column:hide_children_in_menu;comment:在菜单中隐藏下级"`
	HideInBreadcrumb   *bool          `json:"hideInBreadcrumb" gorm:"column:hide_in_breadcrumb;comment:在面包屑中隐藏"`
	HideInMenu         *bool          `json:"hideInMenu" gorm:"column:hide_in_menu;comment:在菜单中隐藏"`
	HideInTab          *bool          `json:"hideInTab" gorm:"column:hide_in_tab;comment:在标签栏中隐藏"`
	Icon               string         `json:"icon" gorm:"column:icon;comment:菜单图标"`
	IframeSrc          string         `json:"iframeSrc" gorm:"column:iframe_src;comment:内嵌Iframe的URL"`
	KeepAlive          *bool          `json:"keepAlive" gorm:"column:keep_alive;comment:是否缓存页面"`
	Link               string         `json:"link" gorm:"column:link;comment:外链页面的URL"`
	MaxNumOfOpenTab    *int           `json:"maxNumOfOpenTab" gorm:"column:max_num_of_open_tab;comment:同一个路由最大打开的标签数"`
	NoBasicLayout      *bool          `json:"noBasicLayout" gorm:"column:no_basic_layout;comment:无需基础布局"`
	OpenInNewWindow    *bool          `json:"openInNewWindow" gorm:"column:open_in_new_window;comment:是否在新窗口打开"`
	Order              *int           `json:"order" gorm:"column:menu_order;comment:菜单排序"`
	Query              datatypes.JSON `json:"query" gorm:"column:query;comment:额外的路由参数"`
	Title              string         `json:"title" gorm:"column:title;comment:菜单标题"`
}

type SysBaseMenuParameter struct {
	global.BASE_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}
