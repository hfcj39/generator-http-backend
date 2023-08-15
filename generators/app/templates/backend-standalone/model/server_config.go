package model

import (
	"<%= displayName %>/global"
	"gorm.io/datatypes"
)

// ServerConfig
// Type string json:"type"
// Brand string json:"brand"
// Project string json:"project
// Region string `json:"region"
// Location string json:"location"
// OS string json:"os"
// MLU string json:"mlu"
type ServerConfig struct {
	global.BASE_MODEL
	ConfigName  string         `json:"config_name" gorm:"comment:配置名称;unique;not null"`
	Description string         `json:"description"`
	ConfigType  string         `json:"config_type" gorm:"comment:配置类型,bool,number,array,string;not null"`
	ConfigScope string         `json:"config_scope" gorm:"comment:配置作用域;not null;default:'server'"`
	ConfigValue datatypes.JSON `json:"config_value"`
}
