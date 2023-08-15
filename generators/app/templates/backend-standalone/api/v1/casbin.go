package v1

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model/request"
	"<%= displayName %>/model/response"
	"<%= displayName %>/service"
	"<%= displayName %>/utils/e"

	"github.com/gin-gonic/gin"
)

// GetCasbinList
// @Summary GetCasbinList
// @Tags Casbin
// @Description 获取casbin规则
// @ID GetCasbinList
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Router /casbin/list [get]
func GetCasbinList(c *gin.Context) {
	rst := service.GetCasbinRuleList()
	response.OkWithData(rst, c)
}

func AddCasbinRole(c *gin.Context) {}

// 添加的接口都在source中添加，暂时不提供界面添加接口功能

// UpdateCasbinRule
// @Summary UpdateCasbinRule
// @Tags Casbin
// @Description 修改casbin规则
// @ID UpdateCasbinRule
// @Accept  json
// @Produce  json
// @Param oldPolicy body []string true "旧规则"
// @Param newPolicy body []string true "新规则"
// @Success 200 {object} response.Response
// @Router /casbin/update [post]
func UpdateCasbinRule(c *gin.Context) {
	p := &request.UpdateCasbin{}
	if err := c.ShouldBindJSON(p); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	rst, err := service.UpdateCasbin(p.OldPolicy, p.NewPolicy)
	if err != nil {
		response.FailWithDetailed(7, err.Error(), "操作失败", c)
		return
	}
	if rst {
		response.OkWithData(rst, c)
	} else {
		response.FailWithMessage("未修改", c)
	}

}

// UpdateCasbinRuleById
// @Summary UpdateCasbinRuleById
// @Tags Casbin
// @Description 通过id修改v0
// @ID UpdateCasbinRuleById
// @Accept  json
// @Produce  json
// @Param id body int true "id"
// @Param v0 body string true "v0"
// @Param v0 body string true "v2"
// @Success 200 {object} response.Response
// @Router /casbin/updateById [post]
func UpdateCasbinRuleById(c *gin.Context) {
	p := &request.UpdateCasbinById{}
	if err := c.ShouldBindJSON(p); err != nil {
		global.LOG.Error(err.Error())
		response.FailWithDetailed(e.InvalidParams, err.Error(), e.GetMsg(e.InvalidParams), c)
		return
	}
	err := service.UpdateCasbinById(p.ID, p.V0, p.V2, p.V3)
	if err != nil {
		response.FailWithDetailed(7, err.Error(), "操作失败", c)
		return
	}
	response.Ok(c)
}
