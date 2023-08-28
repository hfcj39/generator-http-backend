package service

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"
	"<%= displayName %>/model/request"
)

//@function: FindByUserName
//@description: 用户名搜索
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func FindByUserName(u *model.User) (err error, userInter *model.User) {
	var user model.User
	err = global.DB.Where("username = ?", u.Username).Preload("Role").First(&user).Error
	return err, &user
}

func FindByUserID(uid uint) (error, *model.User) {
	var user model.User
	err := global.DB.Where("id = ?", uid).Preload("Role").Preload("Authority").First(&user).Error
	return err, &user
}

func OnlyFindUserByID(uid uint) (error, *model.User) {
	var user model.User
	err := global.DB.Where("id = ?", uid).First(&user).Error
	return err, &user
}

func GetUserListByAuthority(authorityId uint, page int, limit int) (error, *[]model.User, *int64) {
	var users []model.User
	var count int64
	err := global.DB.Model(&users).
		Where("authority_id = ?", authorityId).
		Preload("Authority").
		Count(&count).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&users).Error
	return err, &users, &count
}

func GetUserListByRole(roleId uint, page int, limit int) (error, *[]model.User, *int64) {
	var users []model.User
	var count int64
	err := global.DB.Model(&users).
		Where("role_id = ?", roleId).
		Preload("Role").
		Count(&count).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&users).Error
	return err, &users, &count
}

func CreateNewUser(u *model.User) error {
	return global.DB.Create(u).Error
}

func DeleteUser(id uint) error {
	err := global.DB.Unscoped().Delete(&model.User{}, id).Error
	return err
}

func GetUserList(args *request.GetUserListStruct, userId uint) (error, *[]model.User, *int64) {
	users := []model.User{}
	var count int64
	_, user := FindByUserID(userId)
	query := global.DB.Model(&model.User{}).
		Joins("left join roles on roles.id = users.role_id").
		Where("users.username ilike ? or users.display_name ilike ?", "%"+args.Name+"%", "%"+args.Name+"%").
		Where("roles.role_value <= ?", user.Role.RoleValue)
	if args.Role != "" {
		query.Where("roles.role_name ilike ?", "%"+args.Role+"%")
	}
	if user.Role.RoleValue != 777 {
		query.Where("users.id = ?", userId)
	}

	if args.RoleID > 0 {
		query.Where("users.role_id = ?", args.RoleID)
	}

	if args.ExcludedRoleID > 0 {
		query.Where("users.role_id != ? and users.role_id != 1", args.ExcludedRoleID)
	}

	err := query.
		Select([]string{
			"users.id",
			"users.created_at",
			"users.updated_at",
			"users.username",
			"users.display_name",
			"users.custom_group",
			"users.remark",
			"users.header_img",
			"users.role_id",
		}).
		Preload("Role").
		Count(&count).
		Limit(args.Limit).Offset((args.Page - 1) * args.Limit).
		Order("users.updated_at desc").
		Find(&users).Error
	return err, &users, &count
}

func FindUserByID(id uint) (error, *model.User) {
	rst := model.User{}
	err := global.DB.Preload("Role").Preload("Authority").First(&rst, id).Error
	return err, &rst
}

func UpdateUser(u *model.User, args *model.User) (error, *model.User) {
	err := global.DB.Model(u).Select("Remark").Updates(args).Error
	err = global.DB.Model(u).Updates(args).Error
	return err, u
}

func UpdateSelfUserInfo(u *model.User) (error, *model.User) {
	err := global.DB.Model(u).Select("Remark", "HeaderImg", "DisplayName").Updates(u).Error
	return err, u
}

func Login(u *model.User) (error, *model.User) {
	var user model.User
	err := global.DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Role").First(&user).Error
	return err, &user
}

func UpdateUserPassword(userID uint, pw string) error {
	err := global.DB.Model(&model.User{}).Where("id=?", userID).Update("password", pw).Error
	return err
}

func UpdateUsersRole(userIDs []uint, roleID uint) error {
	err := global.DB.Model(&model.User{}).Where("id in ?", userIDs).Update("role_id", roleID).Error
	return err
}
