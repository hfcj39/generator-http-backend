package v1

import (
	"<%= displayName %>/global"
	"<%= displayName %>/middleware"
	"<%= displayName %>/model"
	"<%= displayName %>/model/request"
	"<%= displayName %>/model/response"
	"<%= displayName %>/service"
	"<%= displayName %>/utils/e"

	"github.com/gin-gonic/gin"
)

// CreateRole
// @Summary CreateRole
// @Tags Role
// @Description 添加role
// @ID CreateRole
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Param role_name body string true "角色名"
// @Param role_value body string true "角色值"
// @Param description body string true "描述"
// @Param button_permissions body []int true "按钮权限"
// @Success 200 {object} response.Response
// @Router /role/create [post]
func CreateRole(c *gin.Context) {
	Args := &request.AddRoleStruct{}
	if err := c.ShouldBindJSON(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}

	if Args.ButtonPermissions == nil {
		Args.ButtonPermissions = []int64{}
	}

	role := &model.Role{
		RoleName:    Args.RoleName,
		Description: Args.Description,
		RoleValue:   Args.RoleValue,
		// ButtonPermissions: Args.ButtonPermissions, //todo
	}

	err, rst := service.AddRole(role)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailedMessage("创建角色失败, 请检查是否有同名角色名或角色值", c)
		return
	}
	response.OkDetailed(rst, "新增成功", c)
}

// GetRoleList
// @Summary GetRoleList
// @Tags Role
// @Description 获取Role列表
// @ID GetRoleList
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Param role_name query string false "名称"
// @Param page query number true "page"
// @Param limit query number true "limit"
// @Success 200 {object} response.Response
// @Router /role/list [get]
func GetRoleList(c *gin.Context) {
	Args := &request.GetRoleListStruct{}
	if err := c.ShouldBind(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	claims := c.MustGet("claims").(middleware.RoleClaims)
	err, rst, count := service.GetRoleList(Args, claims.RoleValue)
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

// FindRoleById
// @Summary FindRoleById
// @Tags Role
// @Description 根据ID获取Role
// @ID FindRoleById
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Success 200 {object} response.Response
// @Router /role/detail/:id [get]
func FindRoleById(c *gin.Context) {
	Args := &request.UriIdStruct{}
	if err := c.ShouldBindUri(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, p := service.FindRoleById(Args.ID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("找不到ID所对应的角色记录", c)
		return
	}
	response.OkWithData(p, c)
}

// UpdateRole
// @Summary UpdateRole
// @Tags Role
// @Description 更新Role
// @ID UpdateRole
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Param id body number true "ID"
// @Param role_name body string true "角色名"
// @Param role_value body string true "角色值"
// @Param description body string true "描述"
// @Param button_permissions body []int true "按钮权限"
// @Success 200 {object} response.Response
// @Router /role/update [post]
func UpdateRole(c *gin.Context) {
	Args := &request.UpdateRoleStruct{}
	if err := c.ShouldBindJSON(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, r := service.FindRoleById(Args.ID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("没有找到ID对应的角色记录", c)
		return
	}

	if r.RoleValue == 777 && Args.RoleValue != 777 {
		response.FailWithMessage("禁止修改超级管理员的角色值", c)
		return
	}

	if Args.RoleValue > 777 || Args.RoleValue < 100 {
		response.FailWithMessage("角色值不合法", c)
		return
	}

	r.RoleName = Args.RoleName
	r.RoleValue = Args.RoleValue
	r.Description = Args.Description
	// r.ButtonPermissions = Args.ButtonPermissions // todo：改

	err = service.UpdateRole(r)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailedMessage("更新角色信息失败, 请检查是否有同名角色名或角色值", c)
		return
	}
	response.OkWithData(r, c)
}

// DeleteRole
// @Tags Role
// @Summary DeleteRole
// @Description 删除Role
// @ID DeleteRole
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Success 200 {object} response.Response
// @Router /role/:id [delete]
func DeleteRole(c *gin.Context) {
	Args := &request.UriIdStruct{}
	if err := c.ShouldBindUri(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, r := service.FindRoleById(Args.ID)
	if err != nil {
		response.FailWithMessage("没有找到ID对应的角色记录", c)
		return
	}

	err, _, count := service.GetUserListByRole(Args.ID, 1, 1)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("查询角色关联用户失败", c)
		return
	}

	if *count > 0 {
		response.FailWithDetailedMessage("仍有用户属于此角色,请先修改后再删除", c)
		return
	}

	if err, menus := service.GetMenuByRole(Args.ID); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("查询角色关联菜单失败", c)
		return
	} else if len(menus) > 0 {
		response.FailWithDetailedMessage("仍有菜单属于此角色,请先修改后再删除", c)
		return
	}

	err = service.DeleteRole(r)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("删除角色失败", c)
		return
	}
	response.OkWithMessage("删除角色成功", c)
}

// GetRoleUserList
// @Summary GetRoleUserList
// @Tags Role
// @Description 获取角色下的User列表
// @ID GetUsersByRole
// @Produce  json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param role_id query int true "角色id"
// @Success 200 {object} response.Response
// @Router /role/user/list [get]
func GetRoleUserList(c *gin.Context) {
	Args := &request.GetRoleUserListStruct{}
	if err := c.ShouldBind(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, rst, count := service.GetUserListByRole(Args.RoleID, Args.Page, Args.Limit)
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

// AddRoleUser
// @Summary AddRoleUser
// @Tags Role
// @Description 角色批量添加用户
// @ID AddRoleUser
// @Produce  json
// @Param role_id body int true "角色id"
// @Param user_ids body []int true "用户id列表"
// @Success 200 {object} response.Response
// @Router /role/user/add [post]
func AddRoleUser(c *gin.Context) {
	Args := &request.AddRoleUserStruct{}
	if err := c.ShouldBindJSON(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err := service.UpdateUsersRole(Args.UserIDs, Args.RoleID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器出错", c)
		return
	}
	response.OkWithMessage("成功批量添加角色用户", c)
}

func DeleteRoleUser(c *gin.Context) {
	Args := &request.UriIdStruct{}
	if err := c.ShouldBindUri(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err := service.UpdateUsersRole([]uint{Args.ID}, 3)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器出错", c)
		return
	}
	response.OkWithMessage("成功移除角色用户", c)
}
