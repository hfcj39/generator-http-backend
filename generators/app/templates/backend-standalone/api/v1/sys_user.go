package v1

import (
	"<%= displayName %>/global"
	"<%= displayName %>/middleware"
	"<%= displayName %>/model"
	"<%= displayName %>/model/request"
	"<%= displayName %>/model/response"
	"<%= displayName %>/service"
	"<%= displayName %>/utils"
	"<%= displayName %>/utils/e"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserInfo
// @Tags User
// @Summary 获取用户信息
// @Produce  application/json
// @Success 200 {object} response.Response
// @Router /user/userInfo [get]
func GetUserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	u := model.User{}
	err := global.DB.Preload("Role").First(&u, claims.(middleware.RoleClaims).UserId).Error
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("查询用户信息失败", c)
		return
	}

	type UserPassword struct {
		Password string `json:"password"`
	}
	pw := UserPassword{}

	err = global.DB.Raw(fmt.Sprintf("select password from users where id=%d", claims.(middleware.RoleClaims).UserId)).Scan(&pw).Error
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("查询用户其他信息失败", c)
		return
	}

	var hasPassword bool
	if pw.Password != "" {
		hasPassword = true
	} else {
		hasPassword = false
	}

	response.OkWithData(gin.H{"user_info": u, "has_password": hasPassword}, c)
}

// UpdateUser
// @Summary updateUser
// @Tags User
// @Description 更新User
// @ID updateUser
// @Accept  json
// @Produce  json
// @Param id body int true "ID"
// @Param role_id body int false "角色ID"
// @Param authority_id body int false "Authority ID"
// @Param custom_group body string false "自定义组"
// @Param remark body string false "备注"
// @Param header_img body string false "头像"
// @Param display_name body string false "昵称"
// @Success 200 {object} response.Response
// @Router /user/update [post]
func UpdateUser(c *gin.Context) {
	Args := &request.UpdateUserStruct{}
	if err := c.ShouldBindJSON(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}

	err, user := service.FindUserByID(Args.ID)
	if err == gorm.ErrRecordNotFound {
		response.FailWithMessage("用户不存在", c)
		return
	}

	claims, _ := c.Get("claims")
	opId := claims.(middleware.RoleClaims).UserId
	roleValue := claims.(middleware.RoleClaims).RoleValue
	// 如果要修改他人个人信息,需role高于他
	if opId != user.ID && roleValue <= user.Role.RoleValue {
		response.FailWithMessage("无修改此用户的权限", c)
		return
	}

	if roleValue == 444 {
		_, _ = service.UpdateUser(user, &model.User{
			DisplayName: Args.DisplayName,
			HeaderImg:   Args.HeaderImg,
			Remark:      Args.Remark,
		})
	} else {
		// 修改不了自己的role和authority
		if opId == user.ID {
			Args.RoleID = 0
		}
		if Args.RoleID > 0 {
			err, role := service.FindRoleById(Args.RoleID)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					response.FailWithMessage("角色不存在", c)
					return
				}
				global.LOG.Error(err.Error())
				response.FailWithMessage("服务器出错", c)
				return
			}
			if roleValue <= role.RoleValue {
				response.FailWithMessage("角色配置权限不足", c)
				return
			}
		}
		err, _ = service.UpdateUser(user, &model.User{
			RoleID:      Args.RoleID,
			CustomGroup: Args.CustomGroup,
			Remark:      Args.Remark,
			HeaderImg:   Args.HeaderImg,
			DisplayName: Args.DisplayName,
		})
		if err != nil {
			global.LOG.Error(err.Error())
			response.FailWithMessage("服务器出错", c)
			return
		}
		_, user = service.FindUserByID(Args.ID)
	}
	response.OkDetailed(user, "修改成功", c)
}

// UpdateSelfUserInfo
// @Summary UpdateSelfUserInfo
// @Tags User
// @Description 更新UserInfo
// @ID UpdateSelfUserInfo
// @Accept  json
// @Produce  json
// @Param remark body string false "备注"
// @Param header_img body string false "头像"
// @Param display_name body string false "昵称"
// @Success 200 {object} response.Response
// @Router /user/userInfo/update [post]
func UpdateSelfUserInfo(c *gin.Context) {
	Args := &request.UpdateSelfUserInfoStruct{}
	if err := c.ShouldBindJSON(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}

	claims, _ := c.Get("claims")
	userID := claims.(middleware.RoleClaims).UserId

	err, user := service.FindUserByID(userID)
	if err == gorm.ErrRecordNotFound {
		response.FailWithMessage("用户不存在", c)
		return
	}

	user.Remark = Args.Remark
	user.HeaderImg = Args.HeaderImg
	user.DisplayName = Args.DisplayName

	err, user = service.UpdateSelfUserInfo(user)
	if err != nil {
		response.FailWithMessage("更新个人信息失败", c)
		return
	}

	response.OkDetailed(user, "修改成功", c)
}

// GetUserList
// @Summary getUserList
// @Tags User
// @Description 获取User列表
// @ID getUserList
// @Produce  json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param username query string false "用户名"
// @Param role query string false "角色名"
// @Param authority query string false "权限名"
// @Success 200 {object} response.Response
// @Router /user/list [get]
func GetUserList(c *gin.Context) {
	Args := &request.GetUserListStruct{}
	if err := c.ShouldBind(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	claims, _ := c.Get("claims")
	userId := claims.(middleware.RoleClaims).UserId
	err, rst, count := service.GetUserList(Args, userId)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器出错", c)
		return
	}
	response.OkWithData(gin.H{
		"list":  *rst,
		"total": *count,
	}, c)
}

// FindUserById
// @Summary findUserById
// @Tags User
// @Description 根据ID获取User
// @ID findUserById
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Router /user/detail/:id [get]
func FindUserById(c *gin.Context) {
	Args := &request.UriIdStruct{}
	if err := c.ShouldBindUri(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, p := service.FindUserByID(Args.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.FailWithMessage("用户不存在", c)
			return
		}
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器出错", c)
		return
	}
	response.OkWithData(p, c)
}

// DeleteUser
// @Summary deleteUser
// @Tags User
// @Description 删除User
// @ID deleteUser
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Router /user/:id [delete]
func DeleteUser(c *gin.Context) {
	Args := &request.UriIdStruct{}
	if err := c.ShouldBindUri(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, user := service.FindUserByID(Args.ID)
	if err == gorm.ErrRecordNotFound {
		response.FailWithMessage("用户不存在", c)
		return
	}
	claims, _ := c.Get("claims")
	opRoleValue := claims.(middleware.RoleClaims).RoleValue

	if opRoleValue <= user.Role.RoleValue {
		response.FailWithMessage("权限不足,删除失败", c)
		return
	}
	err = service.DeleteUser(Args.ID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器出错", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// Login
// @Tags User
// @Summary 用户名密码登录
// @Produce  application/json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} response.Response
// @Router /user/login [post]
func Login(c *gin.Context) {
	L := &request.Login{}
	if err := c.ShouldBindJSON(L); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	U := &model.User{Username: L.Username, Password: utils.MD5V(L.Password)}
	// note 默认用户是没有密码的,如果用户登录时密码字段传空字符串,会被request验证器拦截,但这里处理是不够完美的
	err, user := service.Login(U)
	if err != nil {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 生成token
	err, jwtToken := utils.CreateToken(user.Username, user.ID, user.Role.ID)
	if err != nil {
		response.FailWithDetailed(e.ErrorAuthToken, err.Error(), e.GetMsg(e.ErrorAuthToken), c)
		return
	}

	response.OkWithData(gin.H{"username": U.Username, "token": jwtToken}, c)
}

// SetPassword
// @Tags SetPassword
// @Summary 密码设置
// @Produce  application/json
// @Param old_password body string false "老密码"
// @Param password body string true "密码"
// @Param repeat_password body string true "重复密码确认"
// @Success 200 {object} response.Response
// @Router /user/password [post]
func SetPassword(c *gin.Context) {
	Args := &request.SetUserPassword{}
	if err := c.ShouldBindJSON(Args); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	claims, _ := c.Get("claims")
	type UserPassword struct {
		Password string `json:"password"`
	}
	pw := UserPassword{}
	err := global.DB.Raw(fmt.Sprintf("select password from users where id=%d", claims.(middleware.RoleClaims).UserId)).Scan(&pw).Error
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("查询用户信息失败", c)
		return
	}

	msg := "新增密码成功"

	if pw.Password != "" {
		if Args.OldPassword == nil {
			response.FailWithMessage("请输入旧密码进行校验!", c)
			return
		} else {
			if utils.MD5V(*Args.OldPassword) != pw.Password {
				response.FailWithMessage("旧密码错误!", c)
				return
			}
		}
		msg = "修改密码成功"
	}

	_ = service.UpdateUserPassword(claims.(middleware.RoleClaims).UserId, utils.MD5V(Args.Password))
	response.OkWithMessage(msg, c)
}

func Register(c *gin.Context) {
	L := &request.Login{}
	if err := c.ShouldBindJSON(L); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	newUser := &model.User{
		Username:    L.Username,
		Password:    utils.MD5V(L.Password),
		DisplayName: L.Username,

		RoleID: 3, // 444 valid user
	}
	err := service.CreateNewUser(newUser)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("创建新用户失败", c)
		return
	}
	global.LOG.Info("Created new user: " + newUser.Username)
	err, jwtToken := utils.CreateToken(newUser.Username, newUser.ID, newUser.RoleID)
	if err != nil {
		response.FailWithDetailed(e.ErrorAuthToken, err.Error(), e.GetMsg(e.ErrorAuthToken), c)
		return
	}

	response.OkDetailed(gin.H{
		"username": newUser.Username,
		"token":    jwtToken,
	}, "创建成功", c)
}
