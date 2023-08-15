package service

import (
	"<%= displayName %>/config"
	"<%= displayName %>/global"
)

func GetSystemConfig() config.Server {
	return global.CONFIG
}
