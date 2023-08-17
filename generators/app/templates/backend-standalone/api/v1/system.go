package v1

import (
	"TalentQuest/global"
	"TalentQuest/model"
	"TalentQuest/model/response"
	"TalentQuest/service"
	"TalentQuest/utils/e"

	"github.com/gin-gonic/gin"
)

// GetSystemConfig
// @Summary GetSystemConfig
// @Tags System
// @Description 获取平台config信息
// @ID GetSystemConfig
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Router /system/getSystemConfig [get]
func GetSystemConfig(c *gin.Context) {
	config := service.GetSystemConfig()
	response.OkWithData(config, c)
}

// GetSystemVersion
// @Summary GetSystemVersion
// @Tags System
// @Description 获取平台版本
// @ID GetSystemVersion
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /version [get]
func GetSystemVersion(c *gin.Context) {
	rst := model.ServerConfig{}
	err := global.DB.First(&rst, 1).Error
	if err != nil {
		response.FailWithDetailed(e.ERROR, err.Error(), "系统版本获取失败", c)
		return
	}
	response.OkWithData(rst.ConfigValue, c)
}
