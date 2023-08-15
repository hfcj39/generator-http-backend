package service

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"
	"<%= displayName %>/utils"
	"fmt"
	"strings"

	casbinAdapter "github.com/casbin/gorm-adapter/v3"
)

func GetMenuList() (error, []model.SysBaseMenu) {
	menus := []model.SysBaseMenu{}
	// todo: 按照order排序
	global.DB.Preload("Buttons").Find(&menus)
	err, rst := getMenuTree(menus)
	return err, rst
}

func GetMenuByRole(roleId uint) (error, []model.SysBaseMenu) {
	r := new(model.Role)
	var err error
	// todo: 按照order排序
	err = global.DB.Where(roleId).Preload("SysBaseMenus").First(&r).Error
	if err != nil {
		return err, nil
	}
	allMenus := r.SysBaseMenus
	err, rst := getMenuTree(allMenus)
	if rst == nil {
		rst = []model.SysBaseMenu{}
	}
	return err, rst
}

func getMenuTree(allMenus []model.SysBaseMenu) (error, []model.SysBaseMenu) {
	treeMap := make(map[string][]model.SysBaseMenu)
	var err error
	for _, v := range allMenus {
		treeMap[v.ParentName] = append(treeMap[v.ParentName], v)
	}
	menus := treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

func getChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[fmt.Sprint(menu.Name)]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func AddMenu(m model.SysBaseMenu) error {
	err := global.DB.Create(&m).Error
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}
	var r model.Role
	// 使用.where添加出role_id都为0，暂时先用多条语句添加
	err = global.DB.First(&r, 1).Error
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}
	err = global.DB.Model(&r).Association("SysBaseMenus").Append(&m)
	return err
}

func DeleteMenu(id uint) error {
	//假如为父级菜单，删除所有子菜单
	_, menu := GetMenuById(id)
	if menu.ParentName == "0" {
		global.DB.Where("parent_name = ?", menu.Name).Unscoped().Delete(&model.SysBaseMenu{})
	}
	return global.DB.Unscoped().Delete(&model.SysBaseMenu{}, id).Error
}

func UpdateMenu(p model.SysBaseMenu) error {
	return global.DB.Model(&model.SysBaseMenu{}).Where("id = ?", p.ID).Save(p).Error
}

func GetMenuById(id uint) (error, model.SysBaseMenu) {
	var menu model.SysBaseMenu
	err := global.DB.First(&menu, id).Error
	return err, menu
}

func AddRoleMenu(menus []uint, roleId int) error {
	var p model.Role
	// p.SysBaseMenus = menus
	if len(menus) == 0 {
		p.SysBaseMenus = []model.SysBaseMenu{}
	} else {
		global.DB.Find(&p.SysBaseMenus, menus)
	}
	p.ID = uint(roleId)
	var r model.Role
	global.DB.Preload("SysBaseMenus").First(&r, "id = ?", roleId)
	err := global.DB.Model(&r).Association("SysBaseMenus").Replace(&p.SysBaseMenus)
	return err
}

func AddButton(button *model.Button) error {
	return global.DB.Create(button).Error
}

type ButtonAuthFlag struct {
	AuthFlag        string
	IsSameAuthority *bool
}

func GetButtonByRole(roleId uint, roleValue int) (error, *[]model.Button) {
	r := new(model.Role)
	var casbinButton []string
	// role_value->所拥有的接口权限
	var casbinRule []casbinAdapter.CasbinRule
	global.DB.Where("v0::int <= ?", roleValue).Find(&casbinRule)
	for _, rule := range casbinRule {
		if rule.V3 != "-" {
			casbinButton = append(casbinButton, strings.Split(rule.V3, "|")...)
		}
	}

	var buttonList = make([]model.Button, 0)
	// role_id->role拥有的菜单->菜单下的按钮
	err := global.DB.
		Where(roleId).
		Preload("SysBaseMenus.Buttons",
			"buttons.is_active = ? and (buttons.type in ? or buttons.auth_flag in ?)", true, []string{"menu", "token"}, casbinButton).
		First(&r).
		Error

	// 过滤掉没有to_menu权限的button
	var menuNameList = make([]string, 0)
	for _, _m := range r.SysBaseMenus {
		menuNameList = append(menuNameList, _m.Name)
	}
	for _, m := range r.SysBaseMenus {
		for _, b := range m.Buttons {
			// 按钮类型为跳转菜单,但是没有to_menu的权限,在这里剔除
			if b.Type == "menu" && !utils.ItemInStructArray(b.ToMenuName, menuNameList) {
				continue
			}
			if b.Type == "token" && b.RoleValue > roleValue {
				continue
			}
			buttonList = append(buttonList, b)
		}
	}

	return err, &buttonList
}

func DeleteButton(id uint) error {
	err := global.DB.Delete(&model.Button{}, id).Error
	return err
}

func UpdateButton(b *model.Button) error {
	err := global.DB.Select("*").Updates(b).Error
	return err
}

func FindButtonByID(id uint) (error, *model.Button) {
	rst := model.Button{}
	err := global.DB.First(&rst, id).Error
	return err, &rst
}

func GetMenuByName(name string) (error, model.SysBaseMenu) {
	var menu model.SysBaseMenu
	err := global.DB.Model(&model.SysBaseMenu{}).Where("name = ?", name).First(&menu).Error
	return err, menu
}
