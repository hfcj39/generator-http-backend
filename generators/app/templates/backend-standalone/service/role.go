package service

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"
	"<%= displayName %>/model/request"
	"fmt"
)

var operator = map[string]string{
	"gte": ">=",
	"lte": "<=",
	"gt":  ">",
	"lt":  "<",
}

func FindRoleByName(name string) (error, *model.Role) {
	rst := model.Role{}
	err := global.DB.First(&rst, model.Role{
		RoleName: name,
	}).Error
	return err, &rst
}

func AddRole(r *model.Role) (error, *model.Role) {
	err := global.DB.Create(r).Error
	return err, r
}

func FindRoleById(id uint) (error, *model.Role) {
	role := &model.Role{}
	err := global.DB.Model(role).First(role, id).Error
	return err, role
}

func GetRoleList(args *request.GetRoleListStruct, role_value int) (error, *[]model.Role, *int64) {
	roles := []model.Role{}
	var count int64
	DB := global.DB.Model(&model.Role{}).
		Where("role_name ilike ?", "%"+args.RoleName+"%")
	if operator[args.FilterMode] != "" {
		DB = DB.
			Where(fmt.Sprintf("role_value %s %d", operator[args.FilterMode], role_value))
	}
	err := DB.
		Count(&count).
		Limit(args.Limit).Offset((args.Page - 1) * args.Limit).
		Order("updated_at desc").
		Find(&roles).Error
	return err, &roles, &count
}

func UpdateRole(r *model.Role) error {
	err := global.DB.Select("*").Updates(r).Error
	return err
}

func DeleteRole(r *model.Role) error {
	return global.DB.Unscoped().Delete(&model.Role{}, r.ID).Error
}

func GetRoleValueByUserId(id uint) (error, int) {
	var user model.User
	err := global.DB.Preload("Role").First(&user, id).Error
	if err != nil {
		return err, 0
	}
	return err, user.Role.RoleValue
}
