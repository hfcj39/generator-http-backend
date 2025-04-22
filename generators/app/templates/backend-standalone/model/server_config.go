package model

import (
	"<%= displayName %>/global"

	"gorm.io/datatypes"
)

// ServerConfig
type ServerConfig struct {
	global.BASE_MODEL
	ConfigName  string         `json:"configName" gorm:"comment:配置名称;unique;not null"`
	Description string         `json:"description"`
	ConfigType  string         `json:"configType" gorm:"comment:配置类型,bool,number,array,string;not null"`
	ConfigScope string         `json:"configScope" gorm:"comment:配置作用域;not null;default:'server'"`
	ConfigValue datatypes.JSON `json:"configValue"`
}
