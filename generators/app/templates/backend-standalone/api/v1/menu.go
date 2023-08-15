package v1

import (
	"<%= displayName %>/global"
	"<%= displayName %>/middleware"
	"<%= displayName %>/model"
	"<%= displayName %>/model/request"
	"<%= displayName %>/model/response"
	"<%= displayName %>/service"
	"<%= displayName %>/utils/e"
	"strings"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// GetMenu
// @Summary getMenu
// @Tags Menu
// @Description 根据用户获取对应菜单
// @ID getMenu
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Success 200 {object} response.Response 返回路由列表
// @Router /menu/getMenuByUser [post]
func GetMenu(c *gin.Context) {
	claims := c.MustGet("claims").(middleware.RoleClaims)
	_, data := service.FindByUserID(claims.UserId)
	if err, menus := service.GetMenuByRole(data.Role.ID); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器错误", c)
	} else {
		response.OkWithData(menus, c)
	}
}

// GetMenuList
// @Summary getMenuList
// @Tags Menu
// @Description 根据全部菜单，需要su权限登录
// @ID getMenuList
// @Accept  json
// @Produce  json
// @Param Authorization header string true "token"
// @Success 200 {object} response.Response 返回全部路由列表
// @Router /menu/getMenuList [get]
func GetMenuList(c *gin.Context) {
	err, data := service.GetMenuList()
	if err != nil {
		response.FailWithDetailed(7, err.Error(), e.GetMsg(7), c)
		return
	} else {
		response.OkWithData(data, c)
	}
}

// AddMenu
// @Tags Menu
// @Summary AddMenu
// @Description 添加菜单
// @Produce  application/json
// @Param parent_name body string true "父级路由name"
// @Param path body string true "路径"
// @Param name body string true "路由名"
// @Param hide_menu body bool false "是否影藏"
// @Param component body string true "组件路径"
// @Param redirect body string false "跳转路径"
// @Param order_no body int false "排序序号"
// @Param meta body model.Meta true "详见struct,都是vue-router参数"
// @Param Authorization header string true "jwt-token"
// @Success 200 {object} response.Response
// @Router /menu/addMenu [post]
func AddMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err = service.AddMenu(menu)
	if err != nil {
		global.LOG.Error(err.Error())
		if strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique") {
			response.FailWithDetailed(e.ERROR, err.Error(), "path或name值已存在", c)
			return
		}
		response.FailWithDetailed(e.ERROR, err.Error(), "添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// DeleteMenu
// @Tags Menu
// @Summary DeleteMenu
// @Description 删除菜单
// @Produce  application/json
// @Param id body int true "数据库记录主键ID"
// @Param Authorization header string true "jwt-token"
// @Success 200 {object} response.Response
// @Router /menu/deleteMenu [delete]
func DeleteMenu(c *gin.Context) {
	var id request.SingleIdStruct
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err = service.DeleteMenu(id.ID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.ERROR, err.Error(), "操作失败", c)
		return
	}
	response.OkWithData(id, c)
}

// UpdateMenu
// @Tags Menu
// @Summary UpdateMenu
// @Description 更新菜单
// @Produce  application/json
// @Param id body number true "id"
// @Param parent_name body string true "父级路由name"
// @Param path body string true "路径"
// @Param name body string true "路由名"
// @Param hide_menu body bool false "是否影藏"
// @Param component body string true "组件路径"
// @Param redirect body string false "跳转路径"
// @Param order_no body int false "排序序号"
// @Param meta body model.Meta true "详见struct,都是vue-router参数"
// @Param Authorization header string true "jwt-token"
// @Success 200 {object} response.Response
// @Router /menu/updateMenu [post]
func UpdateMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err = service.UpdateMenu(menu)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.ERROR, err.Error(), "操作失败", c)
		return
	}
	response.Ok(c)
}

// GetMenuById
// @Tags Menu
// @Summary GetMenuById
// @Description 根据id获取menu记录
// @Produce  application/json
// @Param id query int true "数据库记录主键ID"
// @Param Authorization header string true "jwt-token"
// @Success 200 {object} response.Response
// @Router /menu/getMenuById [get]
func GetMenuById(c *gin.Context) {
	var id request.SingleIdStruct
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, data := service.GetMenuById(id.ID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.ERROR, err.Error(), "操作失败", c)
		return
	}
	response.OkWithData(data, c)
}

// GetRoleMenu
// @Tags Menu
// @Summary getMenuByRole
// @Description 根据角色id获取menu
// @Produce  application/json
// @Param id query int true "role主键ID"
// @Param Authorization header string true "jwt-token"
// @Success 200 {object} response.Response
// @Router /menu/getMenuByRole [get]
func GetRoleMenu(c *gin.Context) {
	var id request.SingleIdStruct
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	if err, menus := service.GetMenuByRole(uint(id.ID)); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("服务器错误", c)
	} else {
		response.OkWithData(menus, c)
	}
}

// UpdateRoleMenu
// @Tags Menu
// @Summary UpdateRoleMenu
// @Description 更新角色对应菜单，菜单数组需要展开子菜单
// @Produce  application/json
// @Param role_id body int true "角色id"
// @Param menus body []uint true "菜单列表"
// @Param Authorization header string true "jwt-token"
// @Success 200 {object} response.Response
// @Router /menu/UpdateRoleMenu [post]
func UpdateRoleMenu(c *gin.Context) {
	var req request.AddRoleMenu
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err = service.AddRoleMenu(req.Menus, req.RoleId)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.ERROR, err.Error(), "操作失败", c)
		return
	}
	response.OkWithData(req, c)
}

// func DeleteRoleMenu(c *gin.Context) {}

func AddButton(c *gin.Context) {
	Args := &request.AddButton{}
	if err := c.ShouldBindJSON(Args); err != nil {
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}

	err, _ := service.GetMenuByName(Args.SysBaseMenuName)
	if err != nil {
		global.LOG.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			response.FailWithMessage("找不到所传菜单名称对应的记录", c)
			return
		}
		response.FailWithMessage("查询关联菜单失败", c)
		return
	}

	if err = service.AddButton(&model.Button{
		SysBaseMenuName: Args.SysBaseMenuName,
		Description:     Args.Description,
		AuthFlag:        Args.AuthFlag,
		IsActive:        Args.IsActive,
		OrderNo:         *Args.OrderNo,
		IsSameAuthority: Args.IsSameAuthority,
		Type:            Args.Type,
		RoleValue:       Args.RoleValue,
	}); err != nil {
		response.FailWithMessage("创建按钮权限失败", c)
		return
	}

	response.OkWithMessage("添加成功", c)
}

func GetButtonByUser(c *gin.Context) {
	claims := c.MustGet("claims").(middleware.RoleClaims)
	_, userInfo := service.FindByUserID(claims.UserId)
	err, b := service.GetButtonByRole(userInfo.Role.ID, userInfo.Role.RoleValue)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("获取用户按钮权限失败", c)
		return
	}
	response.OkWithData(b, c)
}

func UpdateButton(c *gin.Context) {
	Args := &request.UpdateButton{}
	if err := c.ShouldBindJSON(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err, b := service.FindButtonByID(Args.ID)
	if err == gorm.ErrRecordNotFound {
		global.LOG.Error(err.Error())
		response.FailWithMessage("没有找到ID对应的按钮", c)
		return
	}
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("查询按钮出错", c)
		return
	}

	b.SysBaseMenuName = Args.SysBaseMenuName
	b.AuthFlag = Args.AuthFlag
	b.IsActive = Args.IsActive
	b.OrderNo = *Args.OrderNo
	b.Type = Args.Type
	b.IsSameAuthority = Args.IsSameAuthority
	b.Description = *Args.Description

	err = service.UpdateButton(b)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("更新按钮出错", c)
		return
	}
	response.OkDetailed(b, "修改按钮成功", c)
}

func DeleteButton(c *gin.Context) {
	Args := &request.UriIdStruct{}
	if err := c.ShouldBindUri(Args); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err := service.DeleteButton(Args.ID)
	if err != nil {
		global.LOG.Error(err.Error())
		response.FailWithMessage("删除按钮失败", c)
		return
	}
	response.OkWithMessage("删除按钮成功", c)
}
